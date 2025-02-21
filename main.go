package main

import (
	"project/go-vets-backend/db"
	"project/go-vets-backend/handlers"
)

func main() {
	db.ConnectDB()
	handlers.Handlers()
}
