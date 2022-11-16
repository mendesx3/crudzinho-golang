package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mendesx3/crudzinho-golang/src/router"
)

func init() {
	log.Println("init")
}

func main() {
	log.Println("Starting the application...")
	if godotenv.Load() != nil {
		log.Fatal("error loading .env file")
	}

	setGin()

	log.Println("application UP")
}

func setGin() {
	r := gin.New()
	api := r.Group("")
	gin.ForceConsoleColor()
	r.Use(gin.Logger(), gin.Recovery())
	router.InitRoutes(api)
	r.Run(":5000")
}
