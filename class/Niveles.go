package class

import "pkg/models"

type Niveles struct {
	Niv []models.Nivel
}

func (niv Niveles) Remove() {
	niv.Niv = nil
}

func (niv Niveles) Object() []string {
	var data []string
	for _, valor := range niv.Niv {
		data = append(data, valor.Name)
	}
	return data
}

func NewNiveles(a []models.Nivel) Niveles {
	return Niveles{Niv: a}
}
