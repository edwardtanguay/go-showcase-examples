package main

import (
	"tanguay.info/02320-sqlite/models"
)

func main() {
	a := models.App{}
	a.Port = 9003
	a.Initialize()
	a.Run()
	defer a.DB.Close()
}
