package main

import (
	"mock-http/app"
)

func main() {
	a := app.CreateApp()
	a.Start()
}
