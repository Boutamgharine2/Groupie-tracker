package js

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	js "js/data"
)

func Handlhom(url string, data []js.Info) []js.Info {
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
