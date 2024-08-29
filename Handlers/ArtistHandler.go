package groupie

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	groupie "groupie/data"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	url1 := "https://groupietrackers.herokuapp.com/api/"
	var data groupie.Artist
	if  r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	id := strings.Trim(r.URL.Path, "/artist/")

	num, err := strconv.Atoi(id)
	if err != nil || num <= 0 || num > 52 || num == 21 {
		tmp, err := template.ParseFiles("templete/error.html")
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(404)
		tmp.Execute(w, nil)
		return
	}
	err1 := FetchHandler(url1, &data, strconv.Itoa(num))

	if err1 != nil {
		fmt.Println(err1)
		log.Fatal(err1)
	}
	tmpl := template.Must(template.ParseFiles("templete/general.html"))

	tmpl.Execute(w, data)
}
