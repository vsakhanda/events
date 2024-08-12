package db

import (
	"events_go/models"
)

var Events = []models.Event{
	{Id: 1,
		Name:        "First event",
		TypeId:      1,
		LocationId:  2,
		Capacity:    50,
		Description: "First event",
		PerformerId: 1},
	{Id: 2,
		Name:        "Second event",
		TypeId:      1,
		LocationId:  2,
		Capacity:    50,
		Description: "First event",
		PerformerId: 2},
	{Id: 3,
		Name:        "Third event",
		TypeId:      1,
		LocationId:  2,
		Capacity:    50,
		Description: "First event",
		PerformerId: 3},
}

var Users = []models.User{
	{Id: 1,
		Name:      "Vladimir",
		Surname:   "Monomakh",
		Nickname:  "Monomakh",
		Phone:     "phone",
		Email:     "email",
		Instagram: "instagram",
		Telegram:  "telegram",
		StatusId:  "statusId",
		Blocked:   true},
	{Id: 2,
		Name:      "Viktoria",
		Surname:   "",
		Nickname:  "Peremoga",
		Phone:     "phone",
		Email:     "email",
		Instagram: "instagram",
		Telegram:  "telegram",
		StatusId:  "statusId",
		Blocked:   true},
	{Id: 3,
		Name:      "Samson",
		Surname:   "Solomon",
		Nickname:  "Solomon",
		Phone:     "phone",
		Email:     "email",
		Instagram: "instagram",
		Telegram:  "telegram",
		StatusId:  "statusId",
		Blocked:   true},
}

var Performers = []models.Performer{
	{Id: 1,
		Name:      "Viktoria",
		Surname:   "Peremokha",
		Nickname:  "Viktoria",
		Instagram: "instagram",
		Telegram:  "telegram",
		PhotoLink: ""},
	{Id: 2,
		Name:      "Oksana",
		Surname:   "",
		Nickname:  "SMM star",
		Instagram: "instagram",
		Telegram:  "telegram",
		PhotoLink: ""},
	{Id: 3,
		Name:      "Olya",
		Surname:   "Trainee",
		Nickname:  "Trainee",
		Instagram: "instagram",
		Telegram:  "telegram",
		PhotoLink: ""},
}

var Locations = []models.Location{
	{Name: "Marmelad",
		Capacity:    "15",
		Description: "Nice place",
		Link:        "link",
		Lat:         "lat",
		Lon:         "lon",
	},
	{Name: "Obolon plaza",
		Capacity:    "25",
		Description: "Great hall",
		Link:        "link",
		Lat:         "lat",
		Lon:         "lon",
	},
}
