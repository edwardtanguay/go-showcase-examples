
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "the test worked",
		})
	})
	r.Run(":3355")
}
