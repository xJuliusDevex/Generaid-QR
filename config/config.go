package config

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	AP   fyne.App
	Wind fyne.Window
	Var  fyne.ThemeVariant
}

func Init() Config {
	AP := Config{AP: app.NewWithID("JLS-QR")}
	// AP.AP.Settings().SetTheme() Cambiar de diseño
	AP.Wind = AP.AP.NewWindow("JLS-QR")
	AP.Wind.Resize(fyne.NewSize(580, 400))
	AP.AP.Settings().ThemeVariant()
	image, err := fyne.LoadResourceFromPath("../assets/JLS.png")
	if err != nil {
		log.Fatal(err)
	}
	AP.Wind.SetIcon(image)
	return AP
}
