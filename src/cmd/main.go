package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayocodingit/crud-heroes-go/src/config"
	hero "github.com/ayocodingit/crud-heroes-go/src/modules/heroes/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	db := config.LoadDB()

	hero.Handler(router, db)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
