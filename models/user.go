package models

type User struct {
	//Model

	Id        uint
	Name      string
	Surname   string
	Nickname  string
	Phone     string
	Email     string
	Instagram string
	Telegram  string
	StatusId  string
	Blocked   bool
}
