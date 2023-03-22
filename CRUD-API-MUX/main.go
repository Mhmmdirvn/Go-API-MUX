package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", Users).Methods("GET")
	r.HandleFunc("/users/{id}", User).Methods("GET")
	r.HandleFunc("/users", Create).Methods("POST")
	r.HandleFunc("/users/{id}", Delete).Methods("DELETE")
	r.HandleFunc("/users/{id}", Edit).Methods("PUT")

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}