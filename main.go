package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/karelpelcak/foodApI/db"
	"github.com/karelpelcak/foodApI/router"
)

func main() {
	_ = godotenv.Load()

	client, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	r := router.SetupRouter(client)

	r.Run()
}
