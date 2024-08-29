package groupie

import (
	"encoding/json"
	"net/http"

	groupie "groupie/data"
)

func FetchHandler(url string, data *groupie.Artist, id string) error {
	response, err := http.Get(url + "locations/" + id)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&data.Location)
	response, err = http.Get(url + "dates/" + id)
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
	json.NewDecoder(res.Body).Decode(&data.Rolation)
	return err
}
