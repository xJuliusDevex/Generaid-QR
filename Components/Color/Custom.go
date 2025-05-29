package color

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func (c *CustomTheme) Color(name fyne.ThemeColorName, s fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return Colors.Primary
	case theme.ColorNameButton:
		return color.White
	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}
