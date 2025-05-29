package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	Btn           *widget.Button
	width, height float32
	X, Y          float32
	name          string
	tipo          string
}

func NewButton(name string, a func()) Button {
	Btn := Button{
		Btn:  widget.NewButton(name, a),
		name: name,
		tipo: "button",
	}
	return Btn
}
func NewButton_search(name string, a func()) Button {
	btn := Button{
		Btn:  widget.NewButtonWithIcon(name, theme.Icon(theme.IconNameSearch), a),
		name: name,
		tipo: "search",
	}
	return btn
}

func (btn *Button) Resize(width, height float32) {
	btn.width = width
	btn.height = height
}
func (btn *Button) Move(X, Y float32) {
	btn.X = X
	btn.Y = Y
}
func (btn *Button) Render() *fyne.Container {
	btn.Btn.SetText(btn.name)
	btn.Btn.Resize(fyne.NewSize(btn.width, btn.height))
	btn.Btn.Move(fyne.NewPos(btn.X, btn.Y))
	container := container.NewWithoutLayout(btn.Btn)
	return container
}

func (btn Button) Clone(a func()) Button {
	NewBtn := Button{}
	if btn.tipo == "button" {
		NewBtn = Button{
			Btn:    widget.NewButton(btn.name, a),
			name:   btn.name,
			X:      btn.X,
			Y:      btn.Y,
			width:  btn.width,
			height: btn.height,
		}
	} else {
		if btn.tipo == "search" {
			NewBtn = Button{
				Btn:    widget.NewButtonWithIcon(btn.name, theme.Icon(theme.IconNameSearch), a),
				name:   btn.name,
				X:      btn.X,
				Y:      btn.Y,
				width:  btn.width,
				height: btn.height,
			}
		}
	}
	return NewBtn
}
func (btn *Button) SetName(name string) {
	btn.name = name
}
