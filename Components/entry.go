package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Entry struct {
	Entry         *widget.Entry
	Width, Height float32
	X, Y          float32
}

func NewEntry() Entry {
	Entry := Entry{
		Entry: widget.NewEntry(),
	}
	return Entry
}
func (en *Entry) Placeholder(name string) {
	en.Entry.PlaceHolder = name
}
func (en *Entry) Resize(Width, Height float32) {
	en.Height = Height
	en.Width = Width
}
func (en *Entry) Move(X, Y float32) {
	en.X = X
	en.Y = Y
}
func (en *Entry) Render() *fyne.Container {
	en.Entry.Resize(fyne.NewSize(en.Width, en.Height))
	en.Entry.Move(fyne.NewPos(en.X, en.Y))
	container := container.NewWithoutLayout(en.Entry)
	return container
}
func (en Entry) Clone() Entry {
	Entry := Entry{
		Entry:  widget.NewEntry(),
		X:      en.X,
		Y:      en.Y,
		Width:  en.Width,
		Height: en.Height,
	}
	return Entry
}
