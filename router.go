package main

import "github.com/gin-gonic/gin"

type Handler struct {
	PingModel PingModel
}

type Config struct {
	R         *gin.Engine
	PingModel PingModel
}

func NewHandler(c *Config) *Handler {
	return &Handler{
		PingModel: c.PingModel,
	}
}

func (h *Handler) pingRouter(g *gin.RouterGroup) {
	g.GET("/ping", h.ping())
	g.GET("/pong", h.pong())
	g.GET("/pin", h.pingPongError())
}

func (h *Handler) ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": h.PingModel.getPong(),
		})
	}
}

func (h *Handler) pong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": h.PingModel.getPing(),
		})
	}
}

func (h *Handler) pingPongError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
	}
}
