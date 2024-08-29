package groupie

import (
	"net/http"
	"text/template"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style/" {
		tmp, err := template.ParseFiles("templete/error.html")
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(404)
		tmp.Execute(w, nil)
		return

	}

	fs := http.FileServer(http.Dir("./style/"))
	http.StripPrefix("/style/", (fs)).ServeHTTP(w, r)
}
