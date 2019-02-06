package main

import (
	"fmt"
	hnd "instagram/handlers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080

	fmt.Println("INSTAGRAM")

	for path, handler := range hnd.Manejadores {
		http.HandleFunc(path, handler)
	}

	//http.HandleFunc(hnd.PathInicio, hnd.Lista[hnd.PathInicio])

	//http.HandleFunc("/js/", hnd.Js)
	//http.HandleFunc("/envio", hnd.Insert)
	//http.HandleFunc("/lista", hnd.List)

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	fmt.Println("Servidor abierto en http://localhost:" + strconv.Itoa(port))
}
