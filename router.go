package main

import "github.com/gin-gonic/gin"

func pingRouter(router *gin.Engine) {
	game := router.Group("/game")
	{
		game.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		game.GET("/pong", func(c *gin.Context) {
			c.JSON(400, gin.H{
				"message": "you can't pong first",
			})
		})
	}
}
