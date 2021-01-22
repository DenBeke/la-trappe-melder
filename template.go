package latrappemelder

import (
	"bytes"
	"html/template"
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
