package class

import (
	"fmt"
	"pkg/models"
)

type Idiomas struct {
	Idio []models.Idioma
}

func (idio Idiomas) Remove() {
	idio.Idio = nil
}

func (idio Idiomas) Object() []string {
	var data []string
	for _, valor := range idio.Idio {
		data = append(data, valor.SetName())
	}
	return data
}
func (idio Idiomas) FindID(info string) (int, error) {
	for _, valor := range idio.Idio {
		if valor.SetName() == info {
			return valor.SetId(), nil
		}
	}
	return 0, fmt.Errorf("No se encontro")
}
func NewIdiomas(a []models.Idioma) Idiomas {
	return Idiomas{Idio: a}
}
