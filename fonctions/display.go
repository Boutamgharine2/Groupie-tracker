package js

import (
	"net/http"

	js "js/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var tableau []js.Info

	data := Handlhom(url, tableau)

	Home(w, r, data)
}
