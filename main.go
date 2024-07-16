package main

import (
	"schmo_login_go/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run()
}
