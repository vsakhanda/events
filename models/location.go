package models

type Location struct {
	//Model
	Name        string `json:"name"`
	Capacity    string `json:"capacity"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}
