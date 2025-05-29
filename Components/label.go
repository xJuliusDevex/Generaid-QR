package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Label struct {
	label         *widget.Label
	width, height float32
	X, Y          float32
	name          string
}

func NewLabel() Label {
	label := Label{
		label: widget.NewLabel(" "),
	}
	return label
}
func (l *Label) Resize(width, height float32) {
	l.width = width
	l.height = height
}
func (l *Label) Move(X, Y float32) {
	l.X = X
	l.Y = Y
}
func (l *Label) Render() *fyne.Container {
	l.label.SetText(l.name)
	l.label.Resize(fyne.NewSize(l.width, l.height))
	l.label.Move(fyne.NewPos(l.X, l.Y))
	container := container.NewWithoutLayout(l.label)
	return container
}

func (l Label) Clone() Label {
	NewBtn := Label{
		label:  widget.NewLabel(l.name),
		name:   l.name,
		X:      l.X,
		Y:      l.Y,
		width:  l.width,
		height: l.height,
	}
	return NewBtn
}
func (l *Label) SetName(name string) {
	l.name = name
}
