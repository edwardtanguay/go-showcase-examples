package src

import (
	"bytes"
	"html/template"
)

func Header() string {
	return `
<h1>Northwind Site</h1>
<hr>
<a href="/">Home</a> | <a href="/employees">Employees</a>
<hr>
	`
}

func SiteBegin() string {
	type Data struct {
		SiteTitle string
	}
	data := Data{SiteTitle: Config().SiteTitle}
	t, _ := template.New("page").Parse(
 `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.SiteTitle}}</title>
	<link rel="stylesheet" type="text/css" href="/css/main.css">
</head>
<body>
	`)
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	return tpl.String()
}

func SiteEnd() string {
	return `
</body>
</html>
	`
}
