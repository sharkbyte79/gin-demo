package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sharkbyte79/gin-demo/internal/db"
)

func main() {
	r := gin.Default()

	err := db.ConnectToDatabase()

	if err != nil {
		// Kill the process if connection to postgres failed
		log.Fatalf(`error connecting to database %v`, err)
		os.Exit(1)
	}

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK,
			gin.H{
				"message": "meow meow"})
	})

	r.Run(":8080")
}
