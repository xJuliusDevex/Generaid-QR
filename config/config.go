package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type APP struct {
	Ap   fyne.App
	Wind fyne.Window
}

// type mytheme struct {
// }

func Init() APP {
	//nombre de la empresa
	app := APP{Ap: app.NewWithID("CEID-JLS")}
	app.Ap.Settings().SetTheme(&Theme{})
	app.Wind = app.Ap.NewWindow("CEID-JLS")
	app.Wind.Resize(fyne.NewSize(700, 500)) //tamaño del escritorio}
	app.Wind.SetIcon(ResourceICon)
	return app
}
