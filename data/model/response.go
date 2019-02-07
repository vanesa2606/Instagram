package model

//RLogin struct
type RLogin struct {
	Username   string
	Contrasena string
}

//RIdentificado struct
type RIdentificado struct {
	Identificado bool
}

//RFoto struct
type RFoto struct {
	ID    string
	URL   string
	Texto string
}

//RComentario struct
type RComentario struct {
	Texto  string
	IDFoto string
}
