package main

import (
	"fmt"
	"log"
	"os"
	components "pkg/Components"
	"pkg/class"
	"pkg/config"
	"pkg/db"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/joho/godotenv"
)

var database db.Connect

func main() {
	App := config.Init()
	//Busca env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	database = db.NewConnect(os.Getenv("MY_DB"), os.Getenv("DB"))
	err = database.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	//generalizando
	esp, err := db.GetAllEspecialidad(database)
	if err != nil {
		log.Fatal(err)
	}
	Especialidades := class.NewEspecialidades(esp)
	aux6, err := Especialidades.FindID("Regular")
	if err != nil {
		log.Fatal(err)
	}
	idio2, err := db.GetIdiomas(database, aux6)
	if err != nil {
		log.Fatal(err)
	}
	Idiomas := class.NewIdiomas(idio2)
	aux7, err := Idiomas.FindID("Inglés")
	if err != nil {
		log.Fatal(err)
	}
	nv, err := db.GetNiveles(database, aux7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nv)
	Niveles := class.NewNiveles(nv)
	//datos
	// Especialidad := class.DataAllESP()
	// Idioma := class.DataALLIdio()
	//Label donde solo se muestra
	nombre := components.NewLabel()
	nombre.SetName("Nombre:")
	nombre.Move(5, 5)
	Nro := nombre.Clone()
	Nro.X = 450
	Nro.SetName("Nro:")
	Esp := nombre.Clone()
	Esp.Y = 45
	Esp.SetName("Especialidad:")
	Idio := Esp.Clone()
	Idio.Y = 120
	Idio.SetName("Idioma:")
	nvl := Idio.Clone()
	nvl.SetName("Nivel de Idioma:")
	nvl.Y = 194
	//Donde ingresa los datos
	Ent_name := components.NewEntry()
	Ent_name.Resize(350, 40)
	Ent_name.Move(75, 7)
	Ent_name.Placeholder("Ingrese nombre completo")
	Ent_Nro := Ent_name.Clone()
	Ent_Nro.Width = 50
	Ent_Nro.X = 490
	Ent_Nro.Placeholder("0001")
	//Selectivo
	Sel_esp := components.NewSelect(Especialidades.Object(), func(value string) {

	})
	Sel_esp.Resize(170, 40)
	Sel_esp.Move(12, 80)
	Sel_idi := Sel_esp.Clone(Idiomas.Object(), func(value string) {
	})
	Sel_idi.Y = 154
	Sel_nvl := Sel_idi.Clone(Niveles.Object(), func(s string) {
	})
	Sel_nvl.Y = 228
	//botones
	Btn_esp := components.NewButton_search("Buscar", func() {
		prueba, err := Especialidades.FindID(Sel_esp.Select.Selected)
		if err != nil {
			log.Fatal(err)
		}
		Idiomas.Remove()
		prueba2, err := db.GetIdiomas(database, prueba)
		if err != nil {
			log.Fatal(err)
		}
		Idiomas = class.NewIdiomas(prueba2)
		Sel_idi.Select.SetOptions(Idiomas.Object())
		Sel_idi.Select.Selected = Sel_idi.Select.Options[0]
	})
	Btn_esp.Move(195, 80)
	Btn_esp.Resize(130, 40)
	Btn_idio := Btn_esp.Clone(func() {
		prueba5, err := Idiomas.FindID(Sel_idi.Select.Selected)
		if err != nil {
			log.Fatal(err)
		}
		Niveles.Remove()
		prueba6, err := db.GetNiveles(database, prueba5)
		if err != nil {
			log.Fatal(err)
		}
		Niveles = class.NewNiveles(prueba6)
		Sel_nvl.Select.SetOptions(Niveles.Object())
		Sel_nvl.Select.Selected = Sel_nvl.Select.Options[0]
	})
	Btn_idio.Y = 154
	btn_qr := components.NewButton("Generar QR", func() {
		var QRGen class.CERTIFICADO
		_, err := strconv.Atoi(Ent_Nro.Entry.Text)
		if err != nil {
			dialog.ShowError(fmt.Errorf("No pusiste un numero:%s", Ent_Nro.Entry.Text), App.Wind)
		} else {
			QRGen = class.NewCertificado(Ent_name.Entry.Text, Ent_Nro.Entry.Text, Sel_esp.Select.Selected, Sel_idi.Select.Selected, Sel_nvl.Select.Selected)
			fileDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, _ error) {
				GeneradorQR, err := QRGen.GeneradorQR()
				if err != nil {
					log.Fatal(err)
				}
				uc.Write(GeneradorQR)
			}, App.Wind)

			fileDialog.SetFileName(fmt.Sprintf("%s-%s.png", Ent_Nro.Entry.Text, Ent_name.Entry.Text))
			fileDialog.Show()
		}

	})
	btn_qr.Resize(100, 40)
	btn_qr.Move(420, 326)
	//funciones de dialog

	App.Wind.SetContent(container.NewWithoutLayout(
		nombre.Render(),
		Nro.Render(),
		Ent_name.Render(),
		Ent_Nro.Render(),
		Esp.Render(),
		Sel_esp.Render(),
		Btn_esp.Render(),
		Idio.Render(),
		Sel_idi.Render(),
		Btn_idio.Render(),
		nvl.Render(),
		Sel_nvl.Render(),
		btn_qr.Render(),
	))
	App.Wind.ShowAndRun()
}
