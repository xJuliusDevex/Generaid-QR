package class

import (
	"fmt"
	"pkg/models"
)

type Especialidades struct {
	Esps []models.Especialidad
}

func (esp Especialidades) Object() []string {
	var aux []string
	for _, v := range esp.Esps {
		aux = append(aux, v.SetName())
	}
	return aux
}
func (esp Especialidades) FindID(name string) (int, error) {
	for _, v := range esp.Esps {
		if v.SetName() == name {
			return v.Id_especialidad, nil
		}
	}
	return 0, fmt.Errorf("No se encontro el id de  %s", name)
}
func NewEspecialidades(a []models.Especialidad) Especialidades {
	return Especialidades{
		Esps: a,
	}
}
