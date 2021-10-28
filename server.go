package server

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Welcome to the server")
	})

	http.ListenAndServe(":5000", nil)
}
