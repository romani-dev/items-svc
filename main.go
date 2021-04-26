package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID     int
	Name   string
	Image  string
	UserID int
}

type PathParamTodo struct {
	ID string `uri:"id"`
}

func main() {
	r := gin.Default()
	r.GET("/items", func(c *gin.Context) {

		requestID := c.Request.Header.Get("x-request-id")
		host := c.Request.Header.Get("Host")
		log.Printf("x-request-id: %v | host: %v", requestID, host)
		c.Header("x-request-id", requestID)
		c.JSON(200, []Item{
			{ID: 12, Name: "IPhone", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 13, Name: "Dell XPS", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 14, Name: "Monitor Samsung QLED", UserID: 1, Image: "https://picsum.photos/300/300"},
		})
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		var params PathParamTodo
		c.ShouldBindUri(&params)

		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%v", params.ID)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		var parsedBody map[string]interface{}
		json.Unmarshal(body, &parsedBody)
		fmt.Println(parsedBody)
		c.JSON(200, parsedBody)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
