package latrappemelder

import (
	"strings"
)

var mailParentTemplate = `
<body style="margin: 0;min-height: 100%;font-size: 1.1em;font-family: 'Open Sans'; background: rgb(255,215,0); background: linear-gradient(180deg, rgba(255,215,0,1) 0%, rgba(218,165,32,1) 100%);">


    <div class="container" style="max-width: 600px;padding: 1em;box-sizing: border-box;margin: auto;">

        <div class="card" style="box-sizing: border-box;padding: 1em 1em;background-color: white;border-radius: 0.5em;box-shadow: 30px 30px 60px #b38500,
             -30px -30px 60px #f2b500;">
            {{ .Content }}
        </div><!-- .card -->
        
    </div><!-- .container -->
  
</body>
`

var startupMailTemplate = strings.ReplaceAll(mailParentTemplate, "{{ .Content }}", `
<p>Hey Admin,</p>

<p>La Trappe Melder is starting!</p>

<p>LTM</p>
`)

var signupMailSubject = `Gelieve je emailadres te bevestigen voor de La Trappe Melder`
var signupMailTemplate = strings.ReplaceAll(mailParentTemplate, "{{ .Content }}", `
<p>Dag, {{ .Name }},

<p>
Bedankt om je aan te melden voor de <a href="{{ .AppURL }}" target="_BLANK">La Trappe Melder</a>.<br>
Gelieve <a href="{{ .ConfirmURL }}" target="_BLANK">hier</a> je email adres te bevestigen.
</p>

<p>LTM</p>
`)

var newBatchTemplate = strings.ReplaceAll(mailParentTemplate, "{{ .Content }}", `
<p>Dag, {{ .Name }},

<p>
Goed nieuws! La Trappe Quadrupel Oak Aged Batch #{{ .Batch }} is nu beschikbaar in de <a href="{{ .LaTrappeURL }}" target="_BLANK">online kloosterwinkel</a> van La Trappe!<br>
Laat het smaken!!! 🍻
</p>

<p><a href="{{ .AppURL }}" target="_BLANK">La Trappe Melder</a></p>
<p><br></p>
<p class="notice" style="font-weight: 300;margin-top: 2em;line-height: 1.6em;color: gray;font-size: 0.8em;">Meld je <a href="{{ .UnsubscribeURL }}" target="_BLANK">hier</a> af als je deze mails beu bent.</p>
`)
