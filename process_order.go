package main

import (
	"github.com/gin-gonic/gin"
	rabbit "microservice/MessageBroker"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/process", func(c *gin.Context) {

		order := rabbit.Receive("order")

		if len(order) != 0 {

			c.JSON(http.StatusOK, gin.H{

				"order_detail": order,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{

				"order_detail": "no order doesn't exist",
			})
		}

	})
	r.Run(":6000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
