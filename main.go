package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () { 
	http.HandleFunc("/", func(w http.ResponseWriter, г *http.Request) {
		fmt.Fprint(w, "Hello, Geekbrains")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, г *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
