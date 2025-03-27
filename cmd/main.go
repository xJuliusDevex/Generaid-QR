package main

import (
	"image/png"
	"os"
	"pkg/config"
	"pkg/layouts"
	"pkg/models"
	"pkg/service"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := config.Init()
	//Listas de los modelos
	banderas := models.List_banderas()
	Niveles := models.List_Idiomas()
	Red := models.List_redSocial()
	//Entidades para poner y seleccionar los datos
	Name_entr := widget.NewEntry()
	Nro_entry := widget.NewEntry()
	slctid := widget.NewSelect(models.Set_name(banderas...), func(value string) {
	})
	slctid.SetSelected("Inglés")
	aux := slctid.Selected
	slctNid := widget.NewSelect(models.Set_name_idioma(Niveles, aux), func(s string) {
	})
	slctNid.SetSelected("Basico")
	btn_idi := widget.NewButtonWithIcon("Buscar", theme.Icon(theme.IconNameSearch), func() {
		aux := slctid.Selected
		slctNid.SetSelected("Basico")
		slctNid.SetOptions(models.Set_name_idioma(Niveles, aux))

	})
	Cont_nro := layouts.ContainNro(100, 40, 400, 5, Nro_entry)
	Cont_name := layouts.ContainName(300, 40, 0, 5, Name_entr)
	Cont_idioma := layouts.ContainIdioma(1, 50, 180, 40, slctid, btn_idi)
	Cont_nvl_idio := layouts.Contain_nivel(1, 150, 180, 40, slctNid)
	btn_genered_QR := widget.NewButton("Generar QR", func() {
		if service.Status(Name_entr.Text, Nro_entry.Text) {
			if service.Verify_int(Nro_entry.Text) {
				b := models.Filtrar_bandera(banderas, slctid.Selected)
				Client := models.NewCustomer(service.Mayuscula(Name_entr.Text), b.SetBandera_Qr(), slctNid.Selected, Red, service.Full_space(Nro_entry.Text))
				dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, app.Wind)
						return
					}
					if writer == nil {
						return
					}
					qrCode := Client.GenerarQr()
					filepath := writer.URI().Path()
					if !strings.HasSuffix(filepath, ".png") {
						filepath += ".png"
					}
					file, err := os.Create(filepath)
					if err != nil {
						dialog.ShowError(err, app.Wind)
						return
					}
					_ = writer.Close()
					defer file.Close()
					err = png.Encode(file, qrCode.Image(256))
					if err != nil {
						dialog.ShowError(err, app.Wind)
						return
					}
					dialog.ShowInformation("Éxito", "Éxito el generador QR", app.Wind)
				}, app.Wind)
			} else {
				dialog.ShowInformation("Error", "Por favor, ponga numero", app.Wind)
			}

		} else {
			dialog.ShowInformation("Error", "Por favor completo los datos", app.Wind)
			return
		}
	})
	btn_genered_QR.Resize(fyne.NewSize(140, 40))
	btn_genered_QR.Move(fyne.NewPos(480, 400))
	app.Wind.SetContent(container.NewWithoutLayout(
		Cont_name,
		Cont_nro,
		slctid,
		Cont_idioma,
		Cont_nvl_idio,
		btn_genered_QR))
	app.Wind.ShowAndRun()
}
