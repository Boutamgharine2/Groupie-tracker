package groupie

import (
	"net/http"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style/" {                      
		PageNotFound(w, r)

		return

	}

	fs := http.FileServer(http.Dir("./style/"))               // server les fichiès statique à partir de le repertoire specifique
	http.StripPrefix("/style/", (fs)).ServeHTTP(w, r)
}
