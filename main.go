package main

import (
	"assignment-project-new/database"
	"assignment-project-new/router"
)

var (
	PORT = ":8080"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
