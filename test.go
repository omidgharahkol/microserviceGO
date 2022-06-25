package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		d := c.Param("id")
		f := c.Query("omid")
		f1 := c.Query("omid2")
		fmt.Println(f)
		fmt.Println(f1)
		fmt.Println(d)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
