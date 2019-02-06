package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathPrincipal Ruta raíz
const PathPrincipal string = "/principal"

//PathFoto Ruta raíz
const PathFoto string = "/foto"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de CSS
const PathCSSFiles string = "/css/"

//PathRegistro Ruta de registro de los usuarios
const PathRegistro string = "/registro"

//PathLogin Ruta de login de los usuarios
const PathLogin string = "/login"

//PathclearSession Ruta de login de los usuarios
const PathclearSession string = "/logout"

//PathUploader Ruta para subir archivos
const PathUploader string = "/uploader"

//PathListarFoto Ruta para subir archivos
const PathListarFoto string = "/listarfoto"

//PathComentario Ruta para subir archivos
const PathComentario string = "/comentario"

//PathjsonResponse Ruta de registro de los usuarios
const PathjsonResponse string = "/jsonResponse"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathPrincipal] = PrincipalFile
	Manejadores[PathFoto] = FotoFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CssFile
	Manejadores[PathRegistro] = Registro
	Manejadores[PathLogin] = Login
	Manejadores[PathclearSession] = Logout
	Manejadores[PathUploader] = Uploader
	Manejadores[PathListarFoto] = ListarFoto
	//Manejadores[PathComentario] = Comentario

}
