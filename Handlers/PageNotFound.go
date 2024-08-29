package groupie

import (
	"net/http"
	"text/template"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templete/error.html")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(404)

	err = tmp.Execute(w, nil)
	if err != nil {
		http.Error(w, "serveur error", http.StatusInternalServerError)
		return
	}
}
