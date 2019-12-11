package routers

import (
	"github.com/gin-gonic/gin"

	"goshop/restful/controllers"
)

func Routers() *gin.Engine {
	r := gin.New()
	api := r.Group("/api")
	api.POST("/login", controllers.Login)
	return r
}
