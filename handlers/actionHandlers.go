package handlers

import (
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

//Registro Función que inserta los idiomas en la base de datos local
func Registro(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathRegistro {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var user model.Usuario
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &user)

		fmt.Println(user.Nombre)

		if user.Nombre == "" || user.Username == "" || user.Correo == "" || user.Contrasena == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		user.Contrasena = hashComoCadena

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(user)
		fmt.Fprint(w, string(respuesta))

		go client.Registro(&user)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Login Función que hace el login de la pagina
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLogin {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var user model.Login
		e = json.Unmarshal(bytes, &user)

		if e == nil {
			lista := client.Login(&user)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&lista)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Uploader Función sube archivos
func Uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2000)

	file, fileInto, err := r.FormFile("archivo")

	f, err := os.OpenFile("./files/"+fileInto.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintf(w, fileInto.Filename)
}
