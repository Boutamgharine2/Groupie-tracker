package js

type Info struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}
type ConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relations struct {
	Id             int               `json:"id"`
	DatesLocations map[string]string `json:"datesLocations"`
}
type Final struct {
	Information Info
	Location    Locations
	Dates       ConcertDates
	Rolation    Relations
}
