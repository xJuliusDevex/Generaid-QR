package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Contain_nivel(x, y, w, h float32, slt_nvl_i *widget.Select) *fyne.Container {
	nivel_name := widget.NewLabel("Nivel Idioma:")
	nivel_name.Move(fyne.NewPos(x, y))
	slt_nvl_i.Move(fyne.NewPos(x+5, y+40))
	slt_nvl_i.Resize(fyne.NewSize(w, h))
	conNvlI := container.NewWithoutLayout(slt_nvl_i)
	return container.NewWithoutLayout(nivel_name, conNvlI)
}
