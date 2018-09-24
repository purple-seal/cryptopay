package routes

import "github.com/gin-gonic/gin"

func SetupV1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	SetupInvoice(v1)
}
