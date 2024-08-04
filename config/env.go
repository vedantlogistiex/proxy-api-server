package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, exists := os.LookupEnv("RENDER"); exists {
		log.Println("Running in Render.com environment. Environment variables are set directly.")
		return
	}

	log.Println("Loading .env file..")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
