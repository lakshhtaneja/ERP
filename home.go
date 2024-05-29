package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lakshhtaneja/ERP/api"
)

func main() {
	server := mux.NewRouter()

	server.HandleFunc("/student/details", api.GetStudent).Methods("GET")
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
