package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	groupie "groupie/data"
)

func Handler(url string, data []groupie.Band) []groupie.Band {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &data)
	return data
}
