package main

import "tanguay.info/02340-api-mux-modules-sqlite/base"

func main() {
	app := base.App{}
	defer app.DB.Close()
	app.Port = 9003
	app.Initialize()
	app.Run()
}
