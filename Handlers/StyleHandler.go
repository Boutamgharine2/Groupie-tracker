package groupie

import (
	"net/http"
	"strings"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		ErrorHandler(w, r, http.StatusNotFound, "page not found")
		return
	}
	// Serve static files from the "./style/" directory
	fs := http.FileServer(http.Dir("./style/"))
	http.StripPrefix("/style/", (fs)).ServeHTTP(w, r)
}
