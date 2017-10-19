package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Welcome to my world!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy, mam!")
	})
	port := ":1025"
	portenv := os.Getenv("ELENTARI_WORLD_ADDR")
	if portenv != "" {
		port = portenv
	}
	log.Printf("Hola guapa, this is a web server that binds to %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
