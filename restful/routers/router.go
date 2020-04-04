package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//"goshop/restful/controllers"

	_ "goshop/restful/docs"
)

func Routers() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api := r.Group("/api")
	// api.POST("/user/login", controllers.Login)
	// api.POST("/user/register", controllers.Register)
	// api.GET("/user/userlist", controllers.UserList)

	GenRouters(r)
	return r
}
