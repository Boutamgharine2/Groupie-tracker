package groupie

import (
	"encoding/json"
	"net/http"
	groupie "groupie/data"
)

func HomeApi( w http.ResponseWriter, r *http.Request, url string, data []groupie.Band) []groupie.Band {
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	return data
}
