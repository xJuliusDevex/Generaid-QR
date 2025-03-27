package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 1=1 , y=50
// posiciones
func ContainIdioma(x, y, w, h float32, slct_id *widget.Select, btn *widget.Button) *fyne.Container {
	y2 := y + 50
	Idioma_label := widget.NewLabel("Idioma:")

	Idioma_label.Move(fyne.NewPos(float32(x), float32(y)))
	slct_id.Resize(fyne.NewSize(w, h))
	slct_id.Move(fyne.NewPos(float32(x+5), float32(y2)))
	slct_id.Alignment = fyne.TextAlignCenter
	slct_id.SetSelected("Inglés")
	slct_id2 := container.NewWithoutLayout(slct_id)
	btn.Resize(fyne.NewSize(w-80, 40))
	btn.Move(fyne.NewPos(float32(x+200), float32(y2)))
	btn2 := container.NewWithoutLayout(btn)
	Idio_con := container.NewWithoutLayout(Idioma_label)
	Idio_con2 := container.NewWithoutLayout(Idio_con, slct_id2, btn2)
	return Idio_con2
}
