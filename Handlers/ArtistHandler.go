package groupie

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"

	groupie "groupie/data"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) { // traiter les information des artistes dans la second page
	url1 := "https://groupietrackers.herokuapp.com/api/"
	var data groupie.Artist
	var wg sync.WaitGroup

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
	wg.Add(4)

	go func() {
		defer wg.Done()

		FetchHandler(url1+"locations/", &data.Location, strconv.Itoa(num), w, r)
	}()
	go func() {
		defer wg.Done()
		FetchHandler(url1+"dates/", &data.Dates, strconv.Itoa(num), w, r) // remplir les defèrants structures à partir des donnés des APIS
	}()
	go func() {
		defer wg.Done()

		FetchHandler(url1+"artists/", &data.Information, strconv.Itoa(num), w, r)
	}()
	go func() {
		defer wg.Done()
		FetchHandler(url1+"relation/", &data.Rolation, strconv.Itoa(num), w, r)
	}()
	wg.Wait()


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
