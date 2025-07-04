package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageHandler(c *gin.Context) {
	data := gin.H{}

	c.HTML(http.StatusOK, "page.html", data)
}
