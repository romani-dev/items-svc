package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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

var failCount int
var isFail bool
var isSlow bool
var slowTime time.Duration = 1 * time.Second

func main() {
	r := gin.Default()

	r.PUT("/slow", func(c *gin.Context) {
		isSlow = !isSlow

		c.JSON(200, gin.H{
			"isSlow": isSlow,
		})
	})

	r.PUT("/fail", func(c *gin.Context) {
		isFail = !isFail

		c.JSON(200, gin.H{
			"isFail": isFail,
		})
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"is_slow": isSlow,
			"is_fail": isFail,
		})
	})

	r.GET("/items", func(c *gin.Context) {

		if isFail {
			defer func() {
				failCount++
			}()

			if failCount%2 == 0 {
				c.JSON(500, gin.H{
					"error": "Falha na request",
				})
				return
			}
		}

		requestID := c.Request.Header.Get("x-request-id")
		host := c.Request.Header.Get("Host")
		log.Printf("x-request-id: %v | host: %v", requestID, host)
		c.Header("x-request-id", requestID)
		c.JSON(200, []Item{
			{ID: 12, Name: "IPhone", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 13, Name: "Dell XPS", UserID: 1, Image: "https://picsum.photos/300/300"},
			{ID: 14, Name: "Monitor Samsung QLED", UserID: 1, Image: "https://picsum.photos/300/300"},
		})

		if isSlow {
			time.Sleep(slowTime)
		}
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
