package groupie

import (
	"net/http"
	"text/template"

	groupie "groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var tableau []groupie.Band
	url := "https://groupietrackers.herokuapp.com/api/artists"

	if r.URL.Path != "/" || r.URL.Path == "/style" {

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
		return

	}

	if  r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := Handler(url, tableau)
	tmpl := template.Must(template.ParseFiles("templete/index.html"))

	tmpl.Execute(w, data)
}
