package groupie

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	groupie "groupie/data"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) { // traiter les information des artistes dans la second page
	url1 := "https://groupietrackers.herokuapp.com/api/"
	var data groupie.Artist
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	id := strings.Trim(r.URL.Path, "/artist/")

	num, err := strconv.Atoi(id)
	if err != nil || num <= 0 || num > 52 {
		ErrorHandler(w, r, http.StatusNotFound, "page Not found")

		return
	}
	err1 := FetchHandler(url1, &data, strconv.Itoa(num)) // remplir les defèrants structures à partir des donnés des APIS

	if err1 != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	tmpl, err2 := template.ParseFiles("templete/general.html")
	if err2 != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return

	}

	tmpl.Execute(w, data)
}
