package latrappemelder

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/badoux/checkmail"
	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	nekoLog "github.com/neko-neko/echo-logrus/v2/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func (m *LaTrappeMelder) Serve() {

	defer m.r.Close()

	// Validate config
	err := m.config.Validate()
	if err != nil {
		log.Fatalf("Config file is not valid: %v", err)
	}

	// prep echo
	e := echo.New()

	e.HideBanner = true

	// Prometheus metrics
	configMetrics := echoPrometheus.NewConfig()
	configMetrics.Namespace = "latrappemelder"
	e.Use(echoPrometheus.MetricsMiddlewareWithConfig(configMetrics))

	go func() {
		me := echo.New()
		me.HideBanner = true
		me.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		me.Logger.Fatal(me.Start(m.config.MetricsAddress))
	}()

	// request logging
	if m.config.AccessLog {
		nekoLog.Logger().SetOutput(os.Stdout)
		nekoLog.Logger().SetLevel(echoLog.INFO)
		e.Logger = nekoLog.Logger()
		e.Use(logMiddleware.Logger())
	}

	// Rate limitting
	//e.Use(RateLimitMiddleware(1*time.Minute, 5))

	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/subscribe", func(c echo.Context) error {

		u := &struct {
			Name  string
			Email string
		}{}

		if err = c.Bind(u); err != nil {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("Gelieve naam en emailadres op te geven."))
		}

		// Validate
		if u.Email == "" || u.Name == "" {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("Gelieve naam en emailadres op te geven."))
		}
		err := checkmail.ValidateFormat(u.Email)
		if err != nil {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("Gelieve een geldig emailadres op te geven."))
		}

		exists, err := m.r.EmailExists(u.Email)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}
		if exists {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("Dit emailadres is al gekend in onze database."))
		}

		// Insert in DB
		s, err := m.r.Subscribe(u.Name, u.Email)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		// Send confirmation mail
		confirmURL := m.config.AppURL + "/confirm/" + s.UUID
		emailBody, err := htmlStringFromTemplate(signupMailTemplate, struct {
			Name       string
			ConfirmURL string
			AppURL     string
		}{u.Name, confirmURL, m.config.AppURL})
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		err = m.SendMail(s.Email, signupMailSubject, emailBody)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		log.Println("successfully handled new subscription")

		return c.HTML(http.StatusOK, m.simpleHTMLResponse("Registratie voor de La Trappe Melder geslaagd!<br>Check je mailbox om je aanmelding te bevestigen."))

	})

	e.GET("/confirm/:uuid", func(c echo.Context) error {

		uuid := c.Param("uuid")
		if uuid == "" {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("Gelieve een geldige bevestigingslink te gebruiken."))
		}

		_, err := m.r.ConfirmSubscription(uuid)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		log.Println("successfully handled confirmation")

		return c.HTML(http.StatusOK, m.simpleHTMLResponse("Je aanmelding is voltooid! Vanaf nu ontvang je de laatste nieuwe La Trappe Quadrupel Oak Aged batch in je mailbox!"))
	})

	e.GET("/unsubscribe/:email", func(c echo.Context) error {

		email := c.Param("email")
		exists, err := m.r.EmailExists(email)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}
		if !exists {
			return c.HTML(http.StatusBadRequest, m.simpleHTMLResponse("We konden je niet terugvinden in onze database."))
		}

		err = m.r.UnSubscribe(email)
		if err != nil {
			log.Errorln(err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		log.Println("successfully handled unsubscription")

		return c.HTML(http.StatusOK, m.simpleHTMLResponse("Je bent nu afgemeld van de La Trappe melder."))
	})

	e.GET("/", func(c echo.Context) error {

		form, err := htmlStringFromTemplate(formTemplate, m.config)
		if err != nil {
			log.Errorf("couldn't get index: %v", err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		index, err := htmlPageWithContent("La Trappe Melder", form, m.config.HTMLTracking)
		if err != nil {
			log.Errorf("couldn't get index: %v", err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		return c.HTML(http.StatusOK, index)
	})

	// Background job to check for updates
	go func() {
		for {

			m.runMelderJob()

			time.Sleep(15 * time.Minute)

		}
	}()

	log.WithField("config", fmt.Sprintf("%+v", m.config)).Println("Starting La Trappe Melder... 🍻")

	err = m.SendMail(m.config.AdminMail, "La Trappe Melder starting 🍻", startupMailTemplate)
	if err != nil {
		log.Fatalf("couldn't send startup mail: %v", err)
	}

	e.Logger.Fatal(e.Start(m.config.HTTPAddress))

}
