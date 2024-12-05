package routes

import (
	"userpas/controller"

	"github.com/gin-gonic/gin"
)

// Setup Router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	UserC := controller.UserController{}
	User := r.Group("/user")
	{
		User.GET("", UserC.Show)
		User.POST("", UserC.Admin)
	}
	return r
}
