package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//Registro test
func Registro(objeto *model.Usuario) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("nombre: ", objeto.Nombre)

	defer db.Close()
	insert, err := db.Query("INSERT INTO Usuario(Nombre, Username, Correo, Contrasena) VALUES (?, ?, ?, ?)", objeto.Nombre, objeto.Username, objeto.Correo, objeto.Contrasena)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// Login test
func Login(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT Contrasena FROM Usuario WHERE (Username = '" + objeto.Username + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Contrasena FROM Usuario WHERE (Username = '" + objeto.Username + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado string
	for query.Next() {

		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//ConsultaID test
func ConsultaID(username string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID FROM Usuario WHERE (Username = '" + username + "')"

	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (Username = '" + username + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado int
	for query.Next() {

		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//SubirFoto test
func SubirFoto(url string, texto string, id int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Foto(Url, Texto, Usuario_ID) VALUES (?, ?, ?)", url, texto, id)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//MostrarFoto test
func MostrarFoto() []model.RFoto {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID, Url, Texto FROM Foto"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID, Url, Texto FROM Foto")

	if err != nil {
		panic(err.Error())
	}
	resultado := make([]model.RFoto, 0)
	for query.Next() {
		var foto = model.RFoto{}

		err = query.Scan(&foto.ID, &foto.URL, &foto.Texto)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, foto)
	}
	return resultado
}

//Comentario test
func Comentario(objeto model.Comentario, ID int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("id: ", ID)
	fmt.Println("nombre: ", objeto.Texto, "FotoID: ", objeto.ID, "Usuario: ", ID)
	defer db.Close()
	insert, err := db.Query("INSERT INTO Comentario(Texto, Foto_ID, Foto_Usuario_ID) VALUES (?, ?, ?)", objeto.Texto, objeto.ID, ID)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//MostrarComentario test
func MostrarComentario() []model.RComentario {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT Texto, Foto_ID, Foto_Usuario_ID, Username FROM Comentario c INNER JOIN Usuario u ON c.Foto_Usuario_ID=u.ID"
	fmt.Println(comando)
	query, err := db.Query("SELECT Texto, Foto_ID, Foto_Usuario_ID, Username FROM Comentario c INNER JOIN Usuario u ON c.Foto_Usuario_ID=u.ID")

	if err != nil {
		panic(err.Error())
	}
	resultado := make([]model.RComentario, 0)
	for query.Next() {
		var comentario = model.RComentario{}

		err = query.Scan(&comentario.Texto, &comentario.IDFoto, &comentario.IDUsuario, &comentario.Username)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, comentario)
	}
	return resultado
}
