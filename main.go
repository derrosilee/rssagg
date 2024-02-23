package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the Environment variable")
	}
	fmt.Println("PORT", portString)
}
