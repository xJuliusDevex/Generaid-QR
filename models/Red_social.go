package models

type Redes struct {
	Telefono string
	Email    string
	Facebook string
	Sitio    string
}

func List_redSocial() Redes {
	redes := Redes{
		Telefono: "052-58 3000 anexos 3055-3056",
		Email:    "ceid@unjbg.edu.pe",
		Facebook: "facebook.com/@CEIDUNJBG",
		Sitio:    " 🌐 ceid.unjbg.edu.pe",
	}
	return redes
}
