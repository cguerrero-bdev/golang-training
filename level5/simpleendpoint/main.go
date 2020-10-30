package main

import (
	"fmt"
	"log"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}

func handleRequests() {
	http.HandleFunc("/", holaMundo)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {

	handleRequests()
}
