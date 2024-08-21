package main

import (
	"events_go/routes"
)

func main() {

	r := routes.InitRoutes()
	//Запуск сервера
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
