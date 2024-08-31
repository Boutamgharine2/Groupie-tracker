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
		PageError(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	id := strings.Trim(r.URL.Path, "/artist/")

	num, err := strconv.Atoi(id)
	if err != nil || num <= 0 || num > 52 {
		PageError(w, r, http.StatusNotFound, "page Not found")

		return
	}
	go FetchHandler(url1+"locations/", &data.Location, strconv.Itoa(num), w, r)
	FetchHandler(url1+"dates/", &data.Dates, strconv.Itoa(num), w, r) // remplir les defèrants structures à partir des donnés des APIS
	FetchHandler(url1+"artists/", &data.Information, strconv.Itoa(num), w, r)
	FetchHandler(url1+"relation/", &data.Rolation, strconv.Itoa(num), w, r)

	tmpl, err2 := template.ParseFiles("templete/general.html")
	if err2 != nil {
		PageError(w, r, http.StatusInternalServerError, "Internal Server Error")
		return

	}

	tmpl.Execute(w, data)
}
