package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Val int `json:"data"`
}

func main() {
	port := ":8080"
	
	var x int = 0
	fmt.Printf("Starting storage node on port %s...\n", port)

	router := gin.Default()
	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Nett hier, aber waren sie schon mal in Baden-Wuttemberg?",
		})
	})

	router.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"data": x,
		})
	})

	router.POST("/", func(c *gin.Context) {
		var jsonBody Data

		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		x = jsonBody.Val
		c.JSON(200, gin.H{
			"data": x,
		})
	})

	router.DELETE("/", func(c *gin.Context) {
		x = 0
		c.JSON(http.StatusOK, gin.H{
			"data": x,
		})
	})
	router.Run()
}