package latrappemelder

var index = `
<html>
	<head>
		<title>La Trappe Quadrupel Oak Aged Melder</title>
	</head>

	<body>

		<form method="get" action="{{ .AppURL }}/subscribe">
			<div>
				<h3>Subscribe</h3>
				<p><input type="text" name="email" placeholder="E-mail" required /></p>
				<p><input type="text" name="name" placeholder="Name" required /></p>
			
				<p><input type="submit" value="Schrijf je in" /></p>
			</div>
		</form>	


	</body>
</html>
`
