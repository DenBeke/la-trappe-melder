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
