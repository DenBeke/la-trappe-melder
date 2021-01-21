package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetBatchVersion(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't parse html: %w", err)
	}

	title := doc.Find("title").First().Text()

	return getBatchVersionFromPageTitle(title)

}

func getBatchVersionFromPageTitle(title string) (string, error) {

	title = strings.TrimSpace(title)

	titleSplit := strings.Split(title, "#")
	if len(titleSplit) != 2 {
		return "", fmt.Errorf("couldn't parse title didn't find '#' delimiter in title: %q", title)
	}
	batch := titleSplit[1]

	return batch, nil
}
