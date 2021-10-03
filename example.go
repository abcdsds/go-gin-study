package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := &Config{
		R:         r,
		PingModel: &PingPong{"ping", "pong"},
	}

	setupRouter(r, c)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter(r *gin.Engine, c *Config) {
	h := NewHandler(c)
	rgroup := r.Group("/game")
	h.pingRouter(rgroup)
}
