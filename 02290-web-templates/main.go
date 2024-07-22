package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Data struct {
	SiteTitle string
	DaysToLive int
}

func HandleMainRoute (w http.ResponseWriter, r *http.Request) {
	data := Data{"The Info Site", 7}
	t,_ := template.New("page").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.SiteTitle}}</title>
</head>
<body>
	<h1>Info Site</h1>
	<p>Welcome to this site, it will be live in {{.DaysToLive}} days.</p>
</body>
</html>
	`)
	t.Execute(w, data)
}

func main() {

	port := 8827
	
	http.HandleFunc("/", HandleMainRoute)

	fmt.Printf("Listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
