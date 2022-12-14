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

	log.Fatal(http.ListenAndServe(":8085", nil))
}
