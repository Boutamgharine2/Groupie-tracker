package groupie

import (
	"net/http"
	"text/template"

	groupie "groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var tableau []groupie.Band
	url := "https://groupietrackers.herokuapp.com/api/artists"

	if r.URL.Path != "/" {
		PageNotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := Handler(url, tableau)
	tmpl := template.Must(template.ParseFiles("templete/index.html"))

	tmpl.Execute(w, data)
}
