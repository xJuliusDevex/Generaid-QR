package config

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// my theme: Theme (*temas para el escritorio*)
type Theme struct {
}

// Colores personalizado
func (t Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return &color.RGBA{R: 245, G: 245, B: 245, A: 255} // Celeste claro
	case theme.ColorNameButton:
		return &color.RGBA{R: 173, G: 216, B: 230, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Font personaliza las fuentes
func (m Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style) // Usa las fuentes predeterminadas
}

// Icon personaliza los íconos
func (m Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name) // Usa los íconos predeterminados
}

// Size personaliza los tamaños
func (m Theme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 10 // Cambia el padding predeterminado
	case theme.SizeNameText:
		return 16 // Cambia el tamaño del texto
	default:
		return theme.DefaultTheme().Size(name)
	}
}
