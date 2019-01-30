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
func Login(objeto *model.Login) []model.RLogin {
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

	resultado := make([]model.RLogin, 0)
	for query.Next() {
		var user = model.RLogin{}

		err = query.Scan(&user.Contrasena)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, user)
	}
	return resultado
}
