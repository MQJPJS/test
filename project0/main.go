package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Gofer gogogo")

	w.SetContent(widget.NewLabel("Gofer gogogo"))
	w.ShowAndRun()
}
