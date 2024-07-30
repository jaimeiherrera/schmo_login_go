package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h Handlers) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h Handlers) CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}
