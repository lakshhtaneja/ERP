package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ERP"))
		})
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
