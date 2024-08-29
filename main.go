package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie/Handlers"
)

func main() {
	http.HandleFunc("/", groupie.HomeHandler)
	http.HandleFunc("/artist/", groupie.ArtistHandler)
	http.HandleFunc("/style/", groupie.StyleHandler)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
