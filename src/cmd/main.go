package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	http.Handler(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
