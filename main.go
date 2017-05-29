package main

import (
	r "myitcv.io/react"

	"github.com/norunners/helloreact/app"
	"honnef.co/go/js/dom"
)

//go:generate reactGen

func main() {
	document := dom.GetWindow().Document()
	domTarget := document.GetElementByID("app")
	app := app.App()

	r.Render(app, domTarget)
}
