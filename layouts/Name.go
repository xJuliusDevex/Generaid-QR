package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Estandar
// w 330 h 40
// pos  x 5 y 10
func ContainName(w, h, x, y float32, wid_en_name *widget.Entry) *fyne.Container {
	wid_name := widget.NewLabel("Nombre:")
	wid_en_name.PlaceHolder = "Ponga el nombre del estudiante"
	wid_en_name.Resize(fyne.NewSize(w, h))
	WName := container.NewWithoutLayout(wid_en_name)
	contai := container.NewHBox(wid_name, WName)
	contai.Move(fyne.NewPos(x, y))
	return container.NewWithoutLayout(contai)
}
