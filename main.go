package main

import (
	"arnov17/echo-test/db"
	"arnov17/echo-test/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
