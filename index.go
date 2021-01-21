package latrappemelder

import (
	"bytes"
	"fmt"
	"html/template"
)

var index = `
<html>
	<head>
		<title>La Trappe Quadrupel Oak Aged Melder</title>
	</head>

	<body>

		<form method="post" action="{{ .ListmonkURL }}/subscription/form" class="listmonk-form">
			<div>
				<h3>Subscribe</h3>
				<p><input type="text" name="email" placeholder="E-mail" required /></p>
				<p><input type="text" name="name" placeholder="Name" required /></p>
			
				<p>
				<input id="a6458" type="checkbox" name="l" value="{{ .ListmonkListID }}" />
				<label for="a6458">la-trappe-melder</label>
				</p>
				<p><input type="submit" value="Subscribe" /></p>
			</div>
		</form>	


	</body>
</html>
`

func GetIndex(data interface{}) (string, error) {

	t, err := template.New("").Parse(index)
	if err != nil {
		return "", fmt.Errorf("couldn't parse index template: %w", err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", fmt.Errorf("couldn't execute index template: %w", err)
	}

	result := tpl.String()

	return result, nil

}
