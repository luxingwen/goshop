package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"goshop/restful/controllers"

	_ "goshop/restful/docs"
)

func Routers() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api := r.Group("/api")
	// api.POST("/user/login", controllers.Login)
	// api.POST("/user/register", controllers.Register)
	// api.GET("/user/userlist", controllers.UserList)

	api := r.Group("/api")
	indexController := controllers.IndexController{}
	{
		api.GET("/index", indexController.Index)
		api.GET("/my_naviga", indexController.MyNaviga)
		api.GET("/index_groom_list/:typ", indexController.GetIndexGroomList)
	}

	store := api.Group("/store")
	storeCategoryController := controllers.StoreCategoryController{}
	storeProductController := controllers.StoreProductController{}
	{
		store.GET("/pid_cate", storeCategoryController.PidByCategory)
		store.GET("/product_category", storeCategoryController.GetProductCategory)
		store.GET("/product_list", storeProductController.ProductList)
		store.GET("/goods_search", storeProductController.GoodsSearch)
	}

	GenRouters(r)
	return r
}
