package db

import (
	"fmt"
	"pkg/models"
)

func GetAllEspecialidad(a Connect) ([]models.Especialidad, error) {
	query := `SELECT * FROM Especialidad`
	row, err := a.Sql.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var Especialidades []models.Especialidad
	for row.Next() {
		var esp models.Especialidad
		if err := row.Scan(&esp.Id_especialidad, &esp.Name); err != nil {
			return nil, err
		}
		Especialidades = append(Especialidades, esp)
	}
	return Especialidades, nil
}
func GetIdiomas(a Connect, id_especialidad int) ([]models.Idioma, error) {
	query := `SELECT I.Id_idioma,I.Id_especialidad,I.Name from idioma as I join especialidad esp ON i.Id_especialidad=esp.Id_especialidad where i.Id_especialidad==?`
	row, err := a.Sql.Query(query, id_especialidad)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var idiomas []models.Idioma
	for row.Next() {
		var i models.Idioma
		err := row.Scan(&i.Id_idioma, &i.Id_especialidad, &i.Name)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		idiomas = append(idiomas, i)
	}

	return idiomas, nil
}
func GetNiveles(a Connect, id_niveles int) ([]models.Nivel, error) {
	query := `SELECT nv.Id_nivel,nv.Id_idioma,nv.Name from nivel as nv join idioma idio ON idio.Id_idioma=nv.Id_idioma where ?==nv.Id_idioma`
	row, err := a.Sql.Query(query, id_niveles)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var Niveles []models.Nivel
	for row.Next() {
		var nv models.Nivel
		err := row.Scan(&nv.Id_nivel, &nv.Id_idioma, &nv.Name)
		if err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		Niveles = append(Niveles, nv)
	}
	return Niveles, nil
}
