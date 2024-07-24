package main

import (
	"tanguay.info/02320-sqlite/base"
)

func main() {
	a := base.App{}
	defer a.DB.Close()
	a.Port = 9003
	a.Initialize()
	a.Run()
}
