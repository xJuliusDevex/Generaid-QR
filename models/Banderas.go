package models

import "fmt"

type Bandera struct {
	Name        string
	Log_bandera string
}

func List_banderas() []Bandera {
	bandera := []Bandera{
		{Name: "Inglés", Log_bandera: "🇺🇸"},
		{Name: "Inglés CEFR", Log_bandera: "🇺🇸"},
		{Name: "Portugués", Log_bandera: "🇧🇷"},
		{Name: "Francés", Log_bandera: "🇫🇷"},
		{Name: "Italiano", Log_bandera: "🇮🇹"},
		{Name: "Alemán", Log_bandera: "🇩🇪"},
		{Name: "Coreano", Log_bandera: "🇰🇷"},
		{Name: "Japonés", Log_bandera: "🇯🇵"},
		{Name: "Chino Mandarín", Log_bandera: "🇨🇳"},
		{Name: "Aymara", Log_bandera: "🇵🇪"},
		{Name: "Quechua", Log_bandera: "🇵🇪"},
	}
	return bandera
}
func Set_name(b ...Bandera) []string {
	names := make([]string, len(b)) // Crear un slice para almacenar los nombres
	for i, v := range b {           // Iterar sobre el slice
		names[i] = v.Name
	}
	return names
}
func Filtrar_bandera(b []Bandera, a string) Bandera {
	for _, bandera := range b {
		if bandera.Name == a {
			return bandera
		}
	}
	return Bandera{}
}
func (a Bandera) SetBandera_Qr() string {
	return fmt.Sprintf("%sCertificado de Idioma - %s%s", a.Log_bandera, a.Name, a.Log_bandera)
}
