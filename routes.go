package main

import (
	"net/http"
	"github.com/gorilla/mux"
) 

type Route struct{
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {//ciclo para iterar las rutas de routes y añadirlas al arreglo de router
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{"PayOrder", "POST", "/pagar_orden", PayOrder},
}