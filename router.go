package main

import "github.com/gin-gonic/gin"

func pingRouter(router *gin.Engine) {
	pingpong := PingModel{}
	game := router.Group("/game")
	{
		game.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": pingpong.pong,
			})
		})

		game.GET("/pong", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": pingpong.ping,
			})
		})

		game.GET("/pin", func(c *gin.Context) {
			c.JSON(400, gin.H{
				"message": "you can only ping, pong",
			})
		})
	}
}
