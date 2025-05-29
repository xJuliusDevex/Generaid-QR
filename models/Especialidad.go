package models

type Especialidad struct {
	Id_especialidad int
	Name            string
}

func (e Especialidad) SetName() string {
	return e.Name
}
func (e Especialidad) SetId() int {
	return e.Id_especialidad
}
func NewEspecialidad(Id_especialidad int, name string) *Especialidad {
	return &Especialidad{
		Id_especialidad: Id_especialidad,
		Name:            name,
	}
}
