package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Couldn't load env file")
	}

	bot_token := os.Getenv("TOKEN")

	log.Println(bot_token)
}
