package main

import (
	"password_warehouse.com/cmd/api"
	"password_warehouse.com/cmd/api/handlers"
)

func main() {

	userHandler := handlers.NewUser()

	api.New(userHandler)
}
