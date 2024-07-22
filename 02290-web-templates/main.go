package main

import (
	"fmt"
	"net/http"
)

func HandleMainRoute (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Info Site</title>
</head>
<body>
	<h1>Info Site</h1>
	<p>Welcome to this site.</p>
</body>
</html>
	`)
}

func main() {

	port := 8827
	
	http.HandleFunc("/", HandleMainRoute)

	fmt.Printf("Listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
