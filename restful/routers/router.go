package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"goshop/restful/common"
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
	api.Use(common.JWTNoMust())
	indexController := controllers.IndexController{}
	loginController := controllers.LoginController{}
	{
		api.GET("/index", indexController.Index)
		api.GET("/logo_url", indexController.GetLogoUrl)
		api.GET("/my_naviga", indexController.MyNaviga)
		api.GET("/index_groom_list/:typ", indexController.GetIndexGroomList)
		api.POST("/login", loginController.Login)
		api.POST("/setCode", loginController.SetCode)
		api.GET("/system_group_data_value/:name", indexController.SystemGroupDataValue)
	}

	store := api.Group("/store")
	storeCategoryController := controllers.StoreCategoryController{}
	storeProductController := controllers.StoreProductController{}
	storeCartController := controllers.StoreCartController{}
	{
		store.GET("/pid_cate", storeCategoryController.PidByCategory)
		store.GET("/product_category", storeCategoryController.GetProductCategory)
		store.GET("/product_list", storeProductController.ProductList)
		store.GET("/hot_products", storeProductController.HotProductList)
		store.GET("/goods_search", storeProductController.GoodsSearch)
		store.GET("/product/:id", storeProductController.Details)
		store.GET("/product_collect/:id", storeProductController.ProductCollect)
		store.GET("/cart_num", storeCartController.GetCartNum)
		store.DELETE("/uncollect_product/:id", storeProductController.UncollectProduct)
		store.POST("/collect_product/:id", storeProductController.CollectProduct)
		store.GET("/user_collect_product", storeProductController.GetUserCollectProduct)
		store.DELETE("/user_collect_product_del/:id", storeProductController.UserCollectProductDel)

	}

	//	 拼团
	pink := api.Group("/pink")

	pinkController := controllers.PinkController{}

	{
		pink.GET("/combination_list", pinkController.CombinationList)
	}

	// 	砍价
	bargain := api.Group("/bargain")

	bargainController := controllers.BargainController{}
	{
		bargain.GET("/bargain_list", bargainController.GetBargainList)
	}

	couponsController := controllers.CouponsController{}

	conpon := api.Group("/coupon")
	{
		conpon.GET("/issue_coupon_list", couponsController.IssueCouponList)
		conpon.GET("/use_coupons/:typ", couponsController.GetUseCoupons)
		conpon.POST("/user_get_coupon/:id", couponsController.UserGetCoupon)
	}

	// GenRouters(r)

	userController := controllers.UserController{}
	userG := api.Group("/user").Use(common.JWT())
	{
		userG.GET("/my", userController.My)
		userG.GET("/my_user_info", userController.MyUserInfo)
		userG.GET("/activity", userController.GetActivity)
		userG.GET("/user_address_list", userController.UserAddressList)
		userG.POST("/edit_user_address", userController.EditUserAddress)
		userG.PUT("/set_user_default_address/:id", userController.SetUserDefaultAddress)
		userG.DELETE("/remove_user_address/:id", userController.RemoveUserAddress)
		userG.GET("/user_address/:id", userController.GetUserAddress)
	}

	return r
}
