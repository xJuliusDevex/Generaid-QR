package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ContainNro(w, h, x, y float32, wid_en_name *widget.Entry) *fyne.Container {
	wid_name := widget.NewLabel("Nro:")
	wid_en_name.PlaceHolder = "Nro"
	wid_en_name.Resize(fyne.NewSize(w, h))
	WName := container.NewWithoutLayout(wid_en_name)
	contai := container.NewHBox(wid_name, WName)
	contai.Move(fyne.NewPos(x, y))
	return container.NewWithoutLayout(contai)

}
