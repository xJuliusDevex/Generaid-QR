package models

type Nivel_idioma struct {
	Idioma string
	Nivel  string
}

func List_Idiomas() []Nivel_idioma {
	idiomas := []Nivel_idioma{
		{Idioma: "Portugués", Nivel: "Basico"},
		{Idioma: "Portugués", Nivel: "Intermedio"},
		{Idioma: "Portugués", Nivel: "Avanzado"},
		{Idioma: "Italiano", Nivel: "Basico"},
		{Idioma: "Italiano", Nivel: "Intermedio"},
		{Idioma: "Italiano", Nivel: "Avanzado"},
		{Idioma: "Alemán", Nivel: "Basico"},
		{Idioma: "Alemán", Nivel: "Intermedio"},
		{Idioma: "Alemán", Nivel: "Avanzado"},
		{Idioma: "Francés", Nivel: "Basico"},
		{Idioma: "Francés", Nivel: "Intermedio"},
		{Idioma: "Francés", Nivel: "Avanzado"},
		{Idioma: "Aymara", Nivel: "Basico"},
		{Idioma: "Aymara", Nivel: "Intermedio"},
		{Idioma: "Aymara", Nivel: "Avanzado"},
		{Idioma: "Quechua", Nivel: "Basico"},
		{Idioma: "Quechua", Nivel: "Intermedio"},
		{Idioma: "Quechua", Nivel: "Avanzado"},
		{Idioma: "Coreano", Nivel: "Basico"},
		{Idioma: "Coreano", Nivel: "Intermedio"},
		{Idioma: "Coreano", Nivel: "Avanzado"},
		{Idioma: "Chino Mandarín", Nivel: "Basico"},
		{Idioma: "Chino Mandarín", Nivel: "Intermedio"},
		{Idioma: "Chino Mandarín", Nivel: "Avanzado"},
		{Idioma: "Inglés", Nivel: "Basico"},
		{Idioma: "Inglés", Nivel: "Intermedio"},
		{Idioma: "Inglés", Nivel: "Intermedio - Avanzado"},
		{Idioma: "Inglés", Nivel: "Avanzado"},
		{Idioma: "Inglés CEFR", Nivel: "Basico A2"},
		{Idioma: "Inglés CEFR", Nivel: "Intermedio B1"},
		{Idioma: "Inglés CEFR", Nivel: "Intermedio - Avanzado B1+"},
		{Idioma: "Inglés CEFR", Nivel: "Avanzado B2"},
		{Idioma: "Japonés", Nivel: "Basico"},
		{Idioma: "Japonés", Nivel: "Intermedio"},
		{Idioma: "Japonés", Nivel: "Intermedio - Avanzado"},
		{Idioma: "Japonés", Nivel: "Avanzado"},
	}
	return idiomas
}
func Set_name_idioma(niv []Nivel_idioma, idioma string) []string {
	var Name []string // Crear un slice para almacenar los nombres
	for _, v := range niv {
		if v.Idioma == idioma {
			Name = append(Name, v.Nivel)
		}
	}
	return Name
}
