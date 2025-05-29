package models

type Nivel struct {
	Id_nivel  int
	Id_idioma int
	Name      string
}

func (nvl Nivel) SetName() string {
	return nvl.Name
}
func (nvl Nivel) SetId_nivel() int {
	return nvl.Id_nivel
}
func (nvl Nivel) SetId_Idioma() int {
	return nvl.Id_idioma
}
func NewNivel(id_nivel, id_idioma int, name string) *Nivel {
	return &Nivel{
		Id_nivel:  id_nivel,
		Id_idioma: id_idioma,
		Name:      name,
	}
}
