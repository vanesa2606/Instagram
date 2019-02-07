package model

//Usuario struct
type Usuario struct {
	Nombre     string
	Username   string
	Correo     string
	Contrasena string
}

//Login struct
type Login struct {
	Username   string
	Contrasena string
}

//Comentario struct
type Comentario struct {
	Texto string
	ID    string
}
