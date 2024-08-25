package js

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	js "js/data"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.URL.Path, "/artist/")
	url1 := "https://groupietrackers.herokuapp.com/api/"
	var data js.Final

	var tableau2 js.ConcertDates
	var tableau3 js.Relations
	err := Traitjs(url1, &data, id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println(t)
	if err1 != nil {
	log.Fatal(err1)
	// }
	// err2, data2 := Traitjs(url2, &tableau2)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// err3, data3 := Traitjs(url3, &tableau3)
	// if err2 != nil {
	// 	log.Fatal(err3)
	// }
	fmt.Println(data)
	General(w, r, data)
	// General(w, r, data2)
	// General(w, r, data3)
}
