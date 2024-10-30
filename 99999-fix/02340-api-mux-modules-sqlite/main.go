package main

import "tanguay.info/02340-api-mux-modules-sqlite/src"

func main() {
	app := src.App{}
	defer app.DB.Close()
	app.Port = 9003
	app.Initialize()
	app.Run()
}
