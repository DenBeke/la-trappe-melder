package latrappemelder

import (
	"bytes"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func htmlStringFromTemplate(htmlTemplate string, data interface{}) (string, error) {

	t, err := template.New("").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	result := tpl.String()

	return result, nil

}

func htmlPageWithContent(title string, content string) (string, error) {

	return htmlStringFromTemplate(index, struct {
		Title   string
		Content string
	}{Title: title, Content: content})

}

func simpleHTMLResponse(content string) string {

	page, err := htmlPageWithContent("La Trappe Melder", "<p>"+content+"</p>")
	if err != nil {
		log.Errorf("Couldn't get html page template: %v", err)
		page = "Oops, something went wrong"
	}

	return page

}
