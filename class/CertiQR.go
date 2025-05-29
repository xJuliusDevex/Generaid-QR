package class

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
)

type CERTIFICADO struct {
	Name           string
	Nro            string
	Especialidades string
	Idioma         string
	Niveles        string
	Redes          string
}

func NewCertificado(Name, Nro, Especialidaddes, Idioma, Niveles string) CERTIFICADO {
	redes := ":D"
	fmt.Println(len(Nro))
	if len(Nro) < 3 {
		for i := len(Nro); i <= 3; i++ {
			Nro = fmt.Sprint("0", Nro)
		}
	}
	return CERTIFICADO{
		Name:           Name,
		Nro:            Nro,
		Especialidades: Especialidaddes,
		Idioma:         Idioma,
		Niveles:        Niveles,
		Redes:          redes,
	}
}

func (a CERTIFICADO) StringFormtaQR() string {
	band := obtenerBanderasIdiomas()
	bandera, err := FindBandera(band, a.Idioma)
	fmt.Println(bandera)
	if err != nil {
		log.Fatal(err)
	}
	year := time.Now().Year()
	yearString := fmt.Sprint(year)

	switch a.Especialidades {
	case "Regular":
		if a.Idioma == "Inglés CEFR" {
			nivel := strings.Split(a.Niveles, " ")
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\nNivel %s: %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, nivel[0], nivel[1], a.Redes)
			return generar
		} else {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\nNivel %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Niveles, a.Redes)
			return generar
		}

	case "Suficiencia":
		if a.Idioma == "Inglés CEFR" {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		} else {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		}

	case "Proficiencia":
		if a.Idioma == "Inglés CEFR" {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		} else {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		}

	case "Maestria":
		if a.Idioma == "Inglés CEFR" {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		} else {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		}

	case "Taller":
		if a.Idioma == "Inglés CEFR" {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		} else {
			generar := fmt.Sprintf("\nN° %s - %s\nNombre:%s\n%sCertificado de Idioma - %s%s\n%s - %s\n%s", a.Nro, yearString, a.Name, bandera, a.Idioma, bandera, a.Especialidades, a.Idioma, a.Redes)
			return generar
		}

	}
	return ""
}
func (a CERTIFICADO) GeneradorQR() ([]byte, error) {
	qr, err := qrcode.New(a.StringFormtaQR(), qrcode.High)
	if err != nil {
		log.Fatal(err)
	}
	return qr.PNG(255)
}

// generador de banderas
func obtenerBanderasIdiomas() map[string]string {
	banderas := map[string]string{
		"Inglés":         "🇺🇸",
		"Inglés CEFR":    "🇺🇸",
		"Portugués":      "🇵🇹",
		"Francés":        "🇫🇷",
		"Italiano":       "🇮🇹",
		"Alemán":         "🇩🇪",
		"Coreano":        "🇰🇷",
		"Japonés":        "🇯🇵",
		"Chino Mandarín": "🇨🇳",
		"Aymara":         "🇵🇪",
		"Quechua":        "🇵🇪",
	}
	return banderas
}
func FindBandera(bandera map[string]string, dato string) (string, error) {
	for key, value := range bandera {
		if key == dato {
			return value, nil
		}
	}
	return "", fmt.Errorf("No se encontro ninguna bandera")
}
