package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("error loading .env file")
	}
	fmt.Println(os.Getenv("PROFILE"))

}
