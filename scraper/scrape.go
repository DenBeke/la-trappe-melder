package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetBatchVersion returns the current batch version from the URL
func GetBatchVersion(url string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "age-gate-nl", Value: "true", HttpOnly: false}) // needed for La Trappe's age gate

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return getBatchFromPage(res.Body)

}

func getBatchFromPage(page io.Reader) (string, error) {

	batches, err := getAllBatchesFromPage(page)
	if err != nil {
		return "", err
	}

	return batches[0], nil

}

const batchesSelector = "h4.clm-list__title"
const prefix = "Batch "

func getAllBatchesFromPage(page io.Reader) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse html: %w", err)
	}

	batchesHTML := doc.Find(batchesSelector)
	if batchesHTML.Length() == 0 {
		return nil, fmt.Errorf("no batches found when parsing")
	}

	batches := make([]string, batchesHTML.Length())

	var batch string
	var errEach error // put this outside since we cannot return directly inside the goquery each

	batchesHTML.EachWithBreak(func(i int, s *goquery.Selection) bool {

		batch, errEach = getBatchFromH4Tag(s.Text())
		if errEach != nil {
			return false
		}

		batches[i] = batch

		return true
	})

	if errEach != nil {
		return nil, fmt.Errorf("couldn't parse batches on pages: %w", errEach)
	}

	return batches, nil
}

func getBatchFromH4Tag(tag string) (string, error) {

	tag = strings.TrimSpace(tag)

	if !strings.HasPrefix(tag, prefix) {
		return "", fmt.Errorf("couldn't parse h4 tag: prefix %q not found", prefix)
	}

	batch := strings.TrimPrefix(tag, prefix)

	_, err := strconv.Atoi(batch)
	if err != nil {
		return "", fmt.Errorf("batch version should be number: got %q", batch)
	}

	return batch, nil

}
