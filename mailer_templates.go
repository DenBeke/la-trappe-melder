package latrappemelder

var startupMailTemplate = `
<p>Hey Admin,</p>

<p>La Trappe Melder is starting!</p>

<p>LTM</p>
`

var signupMailSubject = `Gelieve je emailadres te bevestigen voor de La Trappe Melder`
var signupMailTemplate = `
<p>Dag, {{ .Name }},

<p>
Bedankt om je aan te melden voor de La Trappe Melder.<br>
Gelieve <a href="{{ .ConfirmURL }}" target="_BLANK">hier</a> je email adres te bevestigen.
</p>

<p>LTM</p>
`

var newBatchTemplate = `
<p>Dag, {{ .Name }},

<p>
Goed nieuws! La Trappe Quadrupel Oak Aged Batch #{{ .Batch }} is nu beschikbaar in de <a href="{{ .LaTrappeURL }}" target="_BLANK">online kloosterwinkel</a> van La Trappe!<br>
Laat het smaken!!! 🍻
</p>

<p>LTM</p>
<p><br></p>
<p>Meld je <a href="{{ .UnsubscribeURL }}" target="_BLANK">hier</a> af als je deze mails beu bent.</p>
`
