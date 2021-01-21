package latrappemelder

import (
	"fmt"
	"net/http"
	"os"

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

	e.GET("/", func(c echo.Context) error {

		index, err := GetIndex(m.config)
		if err != nil {
			log.Errorf("couldn't get index: %v", err)
			return c.HTML(http.StatusInternalServerError, "Ooops, something went wrong...")
		}

		return c.HTML(http.StatusOK, index)
	})

	log.WithField("config", fmt.Sprintf("%+v", m.config)).Println("Starting La Trappe Melder... 🍻")

	e.Logger.Fatal(e.Start(m.config.HTTPAddress))

}
