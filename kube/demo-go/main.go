package main

import (
	"context"
	"demo/runtimeid"
	"math/rand/v2"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/get", func(c *gin.Context) {
		id, err := runtimeid.Detect(context.Background(), true)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(200, id)
	})

	router.GET("/magic", func(c *gin.Context) {

		c.JSON(200, map[string]any{"magic": rand.IntN(100001)})
	})
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
