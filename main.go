package main

import (
	"log"
	"net/http"
)

func main() {
	// port number
	port := "8080"
	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	// activate server
	log.Printf("Server listening on port %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
