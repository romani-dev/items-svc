package main

import "github.com/gin-gonic/gin"

type Item struct {
	ID     int
	Name   string
	UserID int
}

func main() {
	r := gin.Default()
	r.GET("/items", func(c *gin.Context) {

		c.JSON(200, []Item{
			{ID: 12, Name: "IPhone", UserID: 1},
			{ID: 13, Name: "Dell XPS", UserID: 1},
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
