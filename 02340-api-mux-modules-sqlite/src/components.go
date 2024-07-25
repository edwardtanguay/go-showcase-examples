package src

func Header() string {
	return `
<h1>Northwind Site</h1>
<hr>
<a href="/">Home</a> | <a href="/employees">Employees</a>
<hr>
	`
}

func SiteBegin() string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>TODO</title>
</head>
<body>
	`
}

func SiteEnd() string {
	return `
</body>
</html>
	`
}
