package main

import (
	"log"
	"os"

	"brand-collab-tracker/config"
	"brand-collab-tracker/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := routes.SetupRouter()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server runnning on http://localhost:%s", port)
	log.Fatal(router.Run(":" + port))
}