package main

import (
	"net/http"

	"github.com/HMPrgm/dynamo-store/internal/storage"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Key string `json:"key"`
	Value []byte `json:"value"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Nett hier, aber waren sie schon mal in Baden-Wuttemberg?",
		})
	})



	store := storage.NewMemoryStore()


	router.POST("/node", func(c *gin.Context) {
		var jsonBody Data

		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// POSSIBLE ERROR: Doesn't check if Post returns error 
		store.Post(jsonBody.Key, jsonBody.Value)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully added key:value pair",
		})
	})
	router.GET("/node", func (c *gin.Context) {

		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no key in query",
			})
			return
		}

		value, err := store.Get(key) 

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"value": value,
		})
	})

	
	// not implemented
	router.DELETE("/node", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "a",
		})
	})
	router.Run()
}