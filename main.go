package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

}
