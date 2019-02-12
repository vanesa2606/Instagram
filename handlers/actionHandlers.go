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

// setSession pone el nombre de usuario proporcionado en un simple mapa de cadena.
//Posteriormente, el manejador de cookies seguro se utiliza para codificar el mapa de valores.
//El valor de sesión resultante (encriptado) se almacena en una http.Cookieinstancia estándar .

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

// La función getUserName implementa toda la secuencia al revés:
// primero, la cookie se lee de la solicitud.
// Luego, el manejador de cookies seguro se utiliza para decodificar / descifrar el valor de la cookie.
// El resultado es un mapa de cadena y se devuelve el nombre de usuario.

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

//clearSession borra la sesión actual configurando una cookie con un negativo MaxAge.
// Posteriormente, la información de la sesión se borra del cliente.

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//Logout función para cerrar la sesión con la ayuda de la función de clearSession
func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
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

			// Coger id de la base de datos. ( Hay que hacer una peticion a la base de datos)

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
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathUploader {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	r.ParseMultipartForm(2000)

	//Coger el archivo y meterlo en una variable
	file, fileInto, err := r.FormFile("archivo")

	//Coger el texto del formulario y merterlo en una variable
	texto := r.FormValue("texto")

	//Coge la funcion getUserName para coger el nombre de usuario
	username := getUserName(r)

	fmt.Println(texto, "Nombre Usuario: ", username)

	f, err := os.OpenFile("./files/"+fileInto.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()
	io.Copy(f, file)

	//La linea de abajo que esta comentada me manda a la página donde está el nombre del archivo
	//fmt.Fprintf(w, fileInto.Filename)

	//Esta linea de aqui abajo me manda a la pagina principal donde están todas las fotos
	http.Redirect(w, r, "/principal", 301)

	//Datos de la base de datos compara el username y le manda el id
	id := client.ConsultaID(username)
	fmt.Println(id)

	//Subir foto a la base de datos
	go client.SubirFoto(fileInto.Filename, texto, id)

}

//ListarFoto Función que devuelve las peticiones de la base de datos dado un filtro
func ListarFoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListarFoto {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	lista := client.MostrarFoto()

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))

}

//Comentario Función parra guardar los comentarios en la base de datos
func Comentario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathComentario {
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
		var comentario model.Comentario
		enTexto := string(bytes)

		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &comentario)
		fmt.Println(comentario.Texto)

		if comentario.Texto == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Funcion para coger el username
		username := getUserName(r)
		//Datos de la base de datos que compara el username y me da el id de ese username
		id := client.ConsultaID(username)
		fmt.Println("IDUsuario: ", id)

		go client.Comentario(comentario, id)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}

}

//ListarComentario Función que devuelve las peticiones de la base de datos dado un filtro
func ListarComentario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListarComentario {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	//Cogemos los datos de la base de datos para listar todos los comentarios
	lista := client.MostrarComentario()

	//Con estas tres lineas lo convierte a json para enviarlo al cliente
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)

	fmt.Fprint(w, string(respuesta))
}
