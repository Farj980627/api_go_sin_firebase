package main

import (
	"net/http"
	"log"
)

func main() {

	router := NewRouter() //Llamada al methodo NewRouter del script routes para traer las rutas y añadirlas al servidor

	server := http.ListenAndServe(":8081", router)//añadir rutas y el puerto, y lanzar el server
	log.Fatal(server)//si algo sale mal muestra el error

	//limpiar cache
	//go clean -cache -modcache -i -r 
}