package main

import (
	"github.com/gin-gonic/gin"
	rabbit "microservice/MessageBroker"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/order/:username", func(c *gin.Context) {
		username := c.Param("username")
		orderType := c.Query("order_type")
		name := c.Query("name")

		orderdetail := map[string]string{"username": username, "group": orderType, "product": name}

		if orderType == "books" || orderType == "cloths" {

			rabbit.Send(orderdetail, "order")
			c.JSON(http.StatusOK, gin.H{
				"result": "The order was registered",
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"result": "the order group doesn't exist",
			})
		}

	})
	r.Run(":4000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
