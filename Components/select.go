package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Select struct {
	Select        *widget.Select
	width, height float32
	X, Y          float32
}

func NewSelect(options []string, onChange func(string)) Select {
	a := Select{
		Select: widget.NewSelect(options, onChange),
	}
	a.Select.Alignment = fyne.TextAlignCenter
	a.Select.Selected = options[0]
	return a
}
func (slt *Select) Resize(width, height float32) {
	slt.width = width
	slt.height = height
}
func (btn *Select) Move(X, Y float32) {
	btn.X = X
	btn.Y = Y
}
func (btn *Select) Render() *fyne.Container {
	btn.Select.Resize(fyne.NewSize(btn.width, btn.height))
	btn.Select.Move(fyne.NewPos(btn.X, btn.Y))
	container := container.NewWithoutLayout(btn.Select)
	return container
}
func (btn Select) Clone(options []string, onChange func(string)) Select {
	NewBtn := Select{}
	NewBtn = Select{
		Select: widget.NewSelect(options, onChange),
		X:      btn.X,
		Y:      btn.Y,
		width:  btn.width,
		height: btn.height}
	NewBtn.Select.Alignment = fyne.TextAlignCenter
	NewBtn.Select.Selected = options[0]
	return NewBtn
}
