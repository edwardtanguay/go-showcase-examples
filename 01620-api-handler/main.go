package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FrameworkRouteHandler(w http.ResponseWriter, r *http.Request) {
	frameworks := []string{"Angular", "React", "Vue"}
	w.Header().Set("Content-Type", "application/json")
	// http.Error(w, "something went wrong", http.StatusInternalServerError)
	json.NewEncoder(w).Encode(frameworks)
}

func main() {

	port := 2346
	
	http.HandleFunc("/frameworks", FrameworkRouteHandler)

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
		<p>frameworks = <a href="http://localhost:%d/frameworks">http://localhost:%d/frameworks</a></p>
</body>
</html>
		`, port, port)
	})

	fmt.Printf("listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
