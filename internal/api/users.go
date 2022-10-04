package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {

	c.JSON(http.StatusOK, "OK")
}
