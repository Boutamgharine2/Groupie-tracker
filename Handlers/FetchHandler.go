package groupie

import (
	"encoding/json"
	"net/http"
)

func FetchHandler(url string, data any, id string, w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(url + id)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(data)
}

/*response, err = http.Get(url + "dates/" + id)
if err != nil {
	return err
}
defer response.Body.Close()
json.NewDecoder(response.Body).Decode(&data.Dates)
response, err = http.Get(url + "artists/" + id)
if err != nil {
	return err
}
defer response.Body.Close()
json.NewDecoder(response.Body).Decode(&data.Information)
res, err1 := http.Get(url + "relation/" + id)
if err1 != nil {
	return err
}
defer response.Body.Close()
json.NewDecoder(res.Body).Decode(&data.Rolation)*/
