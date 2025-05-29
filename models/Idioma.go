package models

type Idioma struct {
	Id_idioma       int
	Id_especialidad int
	Name            string
}

func (i Idioma) SetName() string {
	return i.Name
}
func (i Idioma) SetId() int {
	return i.Id_idioma
}
func (i Idioma) SetId_esp() int {
	return i.Id_especialidad
}
func NewIdioma(id_idioma, id_especialida int, name string) *Idioma {
	return &Idioma{
		Id_idioma:       id_idioma,
		Id_especialidad: id_especialida,
		Name:            name,
	}
}
