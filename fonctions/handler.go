package js

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
)

func Home(w http.ResponseWriter, r *http.Request, data any) {
	if r.URL.Path != "/" { // handel if url was not valide
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	tmpl := template.Must(template.ParseFiles("templete/index.html"))
	// fmt.Println(data)

	tmpl.Execute(w, data)
}

func General(w http.ResponseWriter, r *http.Request, data any) {
	// if r.URL.Path != "/artist/{{.Id}}" { // handel if url was not valide
	// 	http.Error(w, "page not found", http.StatusNotFound)
	// 	return
	// }

	
	fmt.Println(reflect.TypeOf(data))
	// locations, _ := data["locations"].([]string)
	// fmt.Println(locations)

	// // data2 := final{
	// 	Locations: locations,
	// }

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	tmpl := template.Must(template.ParseFiles("templete/general.html"))
	fmt.Println()

	tmpl.Execute(w, data)
}
