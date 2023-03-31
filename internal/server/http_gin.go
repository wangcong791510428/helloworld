package server

import "github.com/gin-gonic/gin"

func RegisterHttpServerGin() *gin.Engine {
	router := gin.Default()
	rootGrp := router.Group("/api")
	{
		userGrp := rootGrp.Group("/user")
		userGrp.GET("/sayhi", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "hello world",
			})
		})
	}

	return router
}
