package main

import (
	"fmt"
	"net/http"
	"github.com/lakshhtaneja/ERP/api"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/student/details", api.GetStudent)
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
