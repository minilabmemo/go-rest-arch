package api

import "github.com/gin-gonic/gin"

func loadRoutes(e *gin.Engine) {
	root := e.Group("service/api/v1")

	user := root.Group("users")
	{
		user.GET("", Users)

	}

}
