package latrappemelder

import (
	"fmt"
	"strconv"

	"github.com/DenBeke/la-trappe-melder/scraper"
	log "github.com/sirupsen/logrus"
)

// runMelderJob will check for new batches:
// if a new batch is found, it will be saved in the register
// and a notification mail is sent to all confirmed subscribers.
func (m *LaTrappeMelder) runMelderJob() {
	b, err := scraper.GetBatchVersion(m.config.LaTrappeURL)
	if err != nil {
		log.Fatalf("couldn't get batch version: %v", err)
	}

	batch, err := strconv.Atoi(b)
	if err != nil {
		log.Fatalf("couldn't atoi batch version: %v", err)
	}

	exists, err := m.r.BatchExists(uint(batch))
	if err != nil {
		log.Fatalf("couldn't lookup batch: %v", err)
	}
	if exists {
		log.Printf("batch %d already exists.", batch)
		return
	}

	// woop, woop! We have a new batch! send them mails!
	log.WithField("batch", batch).Println("Hoeraaaaaa! We found a new La Trappe Batch!")

	// Save batch in database
	err = m.r.AddBatch(uint(batch))
	if err != nil {
		log.Fatalf("couldn't save batch: %v", err)
	}

	// Send mail to all subscribers
	subscribers, err := m.r.GetAllSubscriptions()
	if err != nil {
		log.Fatalf("couldn't get subscribers: %v", err)
	}

	for _, s := range subscribers {

		if !s.Confirmed {
			// skip unverified subscribers
			continue
		}

		mailSubject := fmt.Sprintf("La Trappe Quadrupel Oak Aged Batch #%d is beschikbaar!", batch)
		mailBody, err := htmlStringFromTemplate(newBatchTemplate, struct {
			Name           string
			Batch          int
			UnsubscribeURL string
			LaTrappeURL    string
			AppURL         string
		}{
			Name:           s.Name,
			Batch:          batch,
			UnsubscribeURL: m.config.AppURL + "/unsubscribe/" + s.Email,
			LaTrappeURL:    m.config.LaTrappeURL,
			AppURL:         m.config.AppURL,
		})
		if err != nil {
			log.WithField("subscriber", s).Errorf("couldn't send email to subscriber: %v", err)
			continue
		}

		m.SendMail(s.Email, mailSubject, mailBody)

	}
}
