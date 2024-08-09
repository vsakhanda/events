package main

import (
	"events_go/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	port := 8085
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
