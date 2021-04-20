package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID     int
	Name   string
	Image  string
	UserID int
}

func main() {
	r := gin.Default()
	r.GET("/items", func(c *gin.Context) {

		requestID := c.Request.Header.Get("x-request-id")
		log.Printf("x-request-id: %v", requestID)
		c.Header("x-request-id", requestID)
		c.JSON(200, []Item{
			{ID: 12, Name: "IPhone", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 13, Name: "Dell XPS", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 14, Name: "Monitor Samsung QLED", UserID: 1, Image: "https://picsum.photos/300/300"},
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
