package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Welcome to my world!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy, mam!")
	})
	port := ":1025"
	log.Printf("Hola guapa, this is a web server that binds to %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
