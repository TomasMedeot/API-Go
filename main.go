package main

import (
	"API/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Router
	router := mux.NewRouter().StrictSlash(true)

	//Default
	router.HandleFunc("/", handlers.IndexRoute)

	//Post
	router.HandleFunc("/post", handlers.PostRoute).Methods("POST")

	//Get
	router.HandleFunc("/get", handlers.GetRoute).Methods("GET")

	//Delete
	router.HandleFunc("/delete/{id}", handlers.DeleteRoute).Methods("DELETE")

	//Put
	router.HandleFunc("/put/{id}", handlers.PutRoute).Methods("PUT")

	//Main
	fmt.Println("Server started on port ", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}
