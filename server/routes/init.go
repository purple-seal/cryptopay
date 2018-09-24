package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ps/cryptopay/server/routes/v1"
)

type Routes struct {

}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	root := r.Group("/")

	routes.SetupV1(root)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func init() {

}