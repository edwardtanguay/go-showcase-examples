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
	<b>SITEBEGIN</b>
	` 
}

func SiteEnd() string {
	return `
	<b>SITEEND</b>
	` 
}