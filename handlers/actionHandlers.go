package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//Registro Función que inserta los usuarios en la base de datos local
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

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)

}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//Login Función para acceder a la página
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

	respuesta := false
	if e == nil {
		// datos que recibe del cliente
		var user model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &user)

		fmt.Println(user.Username)

		if user.Username == "" || user.Contrasena == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Contraseña de la base de datos
		password := client.Login(&user)

		// Comprueba que las dos contraseñas sean iguales
		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Contrasena)); err != nil {
			fmt.Printf("No Login")
		} else {
			respuesta = true
			setSession(user.Username, w)
			fmt.Println("Login")

		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, respuesta)
	}

	fmt.Fprintln(w, respuesta)
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
