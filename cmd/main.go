package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	Ap := app.New()
	w := Ap.NewWindow("JLS-QR")
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
