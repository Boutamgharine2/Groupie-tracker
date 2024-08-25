package js

import (
	"encoding/json"
	"net/http"

	js "js/data"
)

func Traitjs(url string, data *js.Final, id string) error {
	response, err := http.Get(url + "location/" + id)
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
	return err
}
