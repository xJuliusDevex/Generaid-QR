package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/skip2/go-qrcode"
)

type Customer struct {
	Name   string
	Idioma string
	Nivel  string
	Red    Redes
	Nro    string
}

func NewCustomer(name, idioma, nivel string, a Redes, nro string) Customer {
	cliente := Customer{
		Name:   name,
		Idioma: idioma,
		Nivel:  nivel,
		Red:    a,
		Nro:    nro,
	}
	return cliente
}
func (a *Customer) GenerarQr() *qrcode.QRCode {
	Contendido := "EMITIDO POR CEID - UNJBG"
	Time := time.Now()
	anoi := strconv.Itoa(Time.Year())
	Contendido = fmt.Sprintf("%s\nN° %s - %s\nNombre:%s\n%s\nNivel %s\nTelefono:%s\nEmail:%s\nFacebook:%s\n%s", Contendido, a.Nro, anoi, a.Name, a.Idioma, a.Nivel, a.Red.Telefono, a.Red.Email, a.Red.Facebook, a.Red.Sitio)
	qr, err := qrcode.New(Contendido, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	return qr
}
