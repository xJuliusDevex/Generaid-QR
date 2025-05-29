package color

import "image/color"

var Colors = struct {
	Primary   color.Color
	Secondary color.Color
	Error     color.Color
}{
	Primary:   color.NRGBA{R: 41, G: 128, B: 185, A: 255}, // Azul
	Secondary: color.NRGBA{R: 25, G: 181, B: 135, A: 255}, // Verde
	Error:     color.NRGBA{R: 200, G: 40, B: 40, A: 255},  // Rojo
}
