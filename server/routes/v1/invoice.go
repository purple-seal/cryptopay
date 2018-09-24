package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ps/cryptopay/server/models"
)

func getList (c *gin.Context) {
	invoice := models.Invoice{}
	invoices := invoice.FindAll()
	c.JSON(http.StatusOK, invoices)
}

func SetupInvoice (router *gin.RouterGroup) {
	invoice := router.Group("/invoices")
	{
		invoice.GET("/", getList)
	}
}
