package main

import (
	"github.com/butlerwang/golang_yaml_rest_api_server/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// SwaggerUI
	router.PathPrefix("/swaggerui").Handler(http.FileServer(http.Dir("./swaggerui/")))

	sub := router.PathPrefix("/api/v1").Subrouter()
	//Create
	sub.Methods("POST").Path("/app").HandlerFunc(handler.CreateApp)
	//Get
	sub.Methods("GET").Path("/app/{name}").HandlerFunc(handler.GetApp)
	//Get all
	sub.Methods("GET").Path("/apps").HandlerFunc(handler.GetApps)
	//Update
	sub.Methods("PUT").Path("/app/{name}").HandlerFunc(handler.UpdateApp)
	//Delete
	sub.Methods("DELETE").Path("/app/{name}").HandlerFunc(handler.DeleteApp)
	//Search
	sub.Methods("GET").Path("/app/search/{name}").HandlerFunc(handler.SearchApp)

	log.Fatal(http.ListenAndServe(":31415", router))
}
