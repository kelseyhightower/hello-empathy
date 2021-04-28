package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Empathy")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}
