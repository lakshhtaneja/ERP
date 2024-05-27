package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ERP"))
		})
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
