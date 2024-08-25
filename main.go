package main

import (
	"fmt"
	"net/http"
	"strings"
	js "js/fonctions"
)

// La structure pour mapper un objet dans la réponse JSON

func main() {
	http.HandleFunc("/", js.Handler)
	http.HandleFunc("/artist/", js.Artist)
	http.HandleFunc("/style/", Style)

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Connected")
		return
	}
}

func Style(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		http.Error(w, "404 : Not Found", http.StatusNotFound)
		return
	}

	stylize := http.FileServer(http.Dir("style"))
	http.StripPrefix("/style", stylize).ServeHTTP(w, r)
}
