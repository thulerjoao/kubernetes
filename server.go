package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func main() {
	http.HandleFunc("/", Hello)

	// Corrigido: porta deve estar no formato ":3000"
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "hello, I'm %s. I'm %s.", name, age)
}
