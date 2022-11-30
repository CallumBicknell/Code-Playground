package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (i IndexController) Home(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
