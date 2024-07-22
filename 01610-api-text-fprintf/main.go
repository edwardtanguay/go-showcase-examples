package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := 4455

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Info API</title>
</head>
<body>
		<h1>Info API</h1>
		<p>Welcome to this API.</p>
</body>
</html>
		`)
	})

	fmt.Printf("Listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
