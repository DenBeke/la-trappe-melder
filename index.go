package latrappemelder

var index = `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>{{ .Title }}</title>
  <meta name="description" content="Mis nooit meer een batch van de La Trappe Quadrupel Oak Aged!">
  <meta name="keywords" content="la trappe, la trappe quadrupel oak aged, bier, trappist" />
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <style type="text/css">

    @import url('https://fonts.googleapis.com/css2?family=Open+Sans:wght@300;400;600;700&display=swap');
    @import url('https://fonts.googleapis.com/css2?family=Kalam:wght@300;400;700&display=swap');

    /*! normalize.css v8.0.1 | MIT License | github.com/necolas/normalize.css */ html{line-height:1.15;-webkit-text-size-adjust:100%}body{margin:0}main{display:block}h1{font-size:2em;margin:.67em 0}hr{box-sizing:content-box;height:0;overflow:visible}pre{font-family:monospace,monospace;font-size:1em}a{background-color:transparent}abbr[title]{border-bottom:none;text-decoration:underline;text-decoration:underline dotted}b,strong{font-weight:bolder}code,kbd,samp{font-family:monospace,monospace;font-size:1em}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:baseline}sub{bottom:-.25em}sup{top:-.5em}img{border-style:none}button,input,optgroup,select,textarea{font-family:inherit;font-size:100%;line-height:1.15;margin:0}button,input{overflow:visible}button,select{text-transform:none}[type=button],[type=reset],[type=submit],button{-webkit-appearance:button}[type=button]::-moz-focus-inner,[type=reset]::-moz-focus-inner,[type=submit]::-moz-focus-inner,button::-moz-focus-inner{border-style:none;padding:0}[type=button]:-moz-focusring,[type=reset]:-moz-focusring,[type=submit]:-moz-focusring,button:-moz-focusring{outline:1px dotted ButtonText}fieldset{padding:.35em .75em .625em}legend{box-sizing:border-box;color:inherit;display:table;max-width:100%;padding:0;white-space:normal}progress{vertical-align:baseline}textarea{overflow:auto}[type=checkbox],[type=radio]{box-sizing:border-box;padding:0}[type=number]::-webkit-inner-spin-button,[type=number]::-webkit-outer-spin-button{height:auto}[type=search]{-webkit-appearance:textfield;outline-offset:-2px}[type=search]::-webkit-search-decoration{-webkit-appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}details{display:block}summary{display:list-item}template{display:none}[hidden]{display:none}

    html {
        min-height:100%;
    }

    body {
        min-height: 100%;
        font-size: 1.1em;
        font-family: 'Open Sans';
        background: rgb(255,215,0);
        background: linear-gradient(180deg, rgba(255,215,0,1) 0%, rgba(218,165,32,1) 100%);

    }

    .container {
        max-width: 800px;
        padding: 2em;
        box-sizing: border-box;

        margin: auto;
    }

    .hero {
        text-align: center;
        padding-top: 1em;
    }

    .hero h1 {
        font-family: 'Kalam';
        font-weight: 700;
        font-size: 3em;
    }

    .hero p {
        font-weight: 400;
    }

    .card {
        box-sizing: border-box;
        padding: 2em;
        background-color: white;

        text-align: center;
        border-radius: 0.5em;

        box-shadow:  30px 30px 60px #b38500,
             -30px -30px 60px #f2b500;  
    }

    p {
        font-weight: 300;
        margin-top: 2em;
        line-height: 1.6em;
    }

    p:first-child {
        margin-top: 0em;
    }

    p.notice {
        color: gray;
        font-size: 0.8em;
    }

    input {
        border-radius: 0.5em;
        outline: none;
        border: 1px solid goldenrod;
        padding: 0.5em;
    }

    input[type="submit"] {
        background-color: goldenrod;
        color: white;
    }

    h1,
    h3 {
        margin: 0;
    }

    a {
        color: inherit;
    }


    footer {
        text-align: center;
        font-size: 0.8em;
    }


  </style>

</head>

<body>


    <div class="container">
        <div class="hero">
            <h1>La Trappe Melder</h1>
            <p>Mis nooit meer een batch van de La Trappe Quadrupel Oak Aged!</p>
        </div><!-- .hero -->
    </div><!-- .container -->


    <div class="container">

        <div class="card">
            {{ .Content }}
        </div><!-- .card -->
        
    </div><!-- .container -->
  

    <footer class="container">
	<p>&copy;2021 <a href="https://denbeke.be" target="_BLANK">Mathias Beke</a> - <a href="https://appoleon.be" target="_BLANK">Appoleon</a> - <a href="https://github.com/DenBeke/la-trappe-melder" target="_BLANK">Github</a></p>
    </footer><!-- .container -->
	
	
	{{ .Tracking }}


</body>
</html>
`

var formTemplate = `
<form method="get" action="{{ .AppURL }}/subscribe">
	<h3>Schrijf je in!</h3>
	<p>Laatste nieuwe batch van de La Trappe Quadrupel Oak Aged in je mailbox?<br>Schrijf je dan hieronder in!</p>
	
	<p><input type="text" name="name" placeholder="Naam" required /></p>
	
	<p><input type="text" name="email" placeholder="E-mail" required /></p>
	

	<p><input type="submit" value="Schrijf je in" /></p>

	<p class="notice">Door je in te schrijven ga je ermee akkoord dat we je emailadres in onze database opslaan en je mails sturen wanneer er nieuwe batches uitkomen. Je kan je te allen tijde uitschrijven.</p>
</form>
`
