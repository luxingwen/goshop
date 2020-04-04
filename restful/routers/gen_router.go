//generate by gen
package routers

import (
	"github.com/gin-gonic/gin"

	"goshop/restful/controllers"
)

func GenRouters(r *gin.Engine) {

	articleController := controllers.ArticleController{}
	articleGroup := r.Group("/articles")
	{
		articleGroup.GET("", articleController.List)
		articleGroup.POST("", articleController.Create)
		articleGroup.DELETE("/:articleId", articleController.Delete)
		articleGroup.PUT("/:articleId", articleController.Put)
		articleGroup.GET("/:articleId", articleController.Get)
		articleGroup.PATCH("/:articleId", articleController.Patch)
	}
	//!!do not delete gen will generate router code at here

	articleCategoryController := controllers.ArticleCategoryController{}
	articleCategoryGroup := r.Group("/articleCategorys")
	{
		articleCategoryGroup.GET("", articleCategoryController.List)
		articleCategoryGroup.POST("", articleCategoryController.Create)
		articleCategoryGroup.DELETE("/:articleCategoryId", articleCategoryController.Delete)
		articleCategoryGroup.PUT("/:articleCategoryId", articleCategoryController.Put)
		articleCategoryGroup.GET("/:articleCategoryId", articleCategoryController.Get)
		articleCategoryGroup.PATCH("/:articleCategoryId", articleCategoryController.Patch)
	}
	//!!do not delete gen will generate router code at here

	expressController := controllers.ExpressController{}
	expressGroup := r.Group("/expresss")
	{
		expressGroup.GET("", expressController.List)
		expressGroup.POST("", expressController.Create)
		expressGroup.DELETE("/:expressId", expressController.Delete)
		expressGroup.PUT("/:expressId", expressController.Put)
		expressGroup.GET("/:expressId", expressController.Get)
		expressGroup.PATCH("/:expressId", expressController.Patch)
	}
	//!!do not delete gen will generate router code at here

	routineAccessTokenController := controllers.RoutineAccessTokenController{}
	routineAccessTokenGroup := r.Group("/routineAccessTokens")
	{
		routineAccessTokenGroup.GET("", routineAccessTokenController.List)
		routineAccessTokenGroup.POST("", routineAccessTokenController.Create)
		routineAccessTokenGroup.DELETE("/:routineAccessTokenId", routineAccessTokenController.Delete)
		routineAccessTokenGroup.PUT("/:routineAccessTokenId", routineAccessTokenController.Put)
		routineAccessTokenGroup.GET("/:routineAccessTokenId", routineAccessTokenController.Get)
		routineAccessTokenGroup.PATCH("/:routineAccessTokenId", routineAccessTokenController.Patch)
	}
	//!!do not delete gen will generate router code at here

	routineFormIdController := controllers.RoutineFormIdController{}
	routineFormIdGroup := r.Group("/routineFormIds")
	{
		routineFormIdGroup.GET("", routineFormIdController.List)
		routineFormIdGroup.POST("", routineFormIdController.Create)
		routineFormIdGroup.DELETE("/:routineFormIdId", routineFormIdController.Delete)
		routineFormIdGroup.PUT("/:routineFormIdId", routineFormIdController.Put)
		routineFormIdGroup.GET("/:routineFormIdId", routineFormIdController.Get)
		routineFormIdGroup.PATCH("/:routineFormIdId", routineFormIdController.Patch)
	}
	//!!do not delete gen will generate router code at here

	routineQrcodeController := controllers.RoutineQrcodeController{}
	routineQrcodeGroup := r.Group("/routineQrcodes")
	{
		routineQrcodeGroup.GET("", routineQrcodeController.List)
		routineQrcodeGroup.POST("", routineQrcodeController.Create)
		routineQrcodeGroup.DELETE("/:routineQrcodeId", routineQrcodeController.Delete)
		routineQrcodeGroup.PUT("/:routineQrcodeId", routineQrcodeController.Put)
		routineQrcodeGroup.GET("/:routineQrcodeId", routineQrcodeController.Get)
		routineQrcodeGroup.PATCH("/:routineQrcodeId", routineQrcodeController.Patch)
	}
	//!!do not delete gen will generate router code at here

	routineTemplateController := controllers.RoutineTemplateController{}
	routineTemplateGroup := r.Group("/routineTemplates")
	{
		routineTemplateGroup.GET("", routineTemplateController.List)
		routineTemplateGroup.POST("", routineTemplateController.Create)
		routineTemplateGroup.DELETE("/:routineTemplateId", routineTemplateController.Delete)
		routineTemplateGroup.PUT("/:routineTemplateId", routineTemplateController.Put)
		routineTemplateGroup.GET("/:routineTemplateId", routineTemplateController.Get)
		routineTemplateGroup.PATCH("/:routineTemplateId", routineTemplateController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeBargainController := controllers.StoreBargainController{}
	storeBargainGroup := r.Group("/storeBargains")
	{
		storeBargainGroup.GET("", storeBargainController.List)
		storeBargainGroup.POST("", storeBargainController.Create)
		storeBargainGroup.DELETE("/:storeBargainId", storeBargainController.Delete)
		storeBargainGroup.PUT("/:storeBargainId", storeBargainController.Put)
		storeBargainGroup.GET("/:storeBargainId", storeBargainController.Get)
		storeBargainGroup.PATCH("/:storeBargainId", storeBargainController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeBargainUserController := controllers.StoreBargainUserController{}
	storeBargainUserGroup := r.Group("/storeBargainUsers")
	{
		storeBargainUserGroup.GET("", storeBargainUserController.List)
		storeBargainUserGroup.POST("", storeBargainUserController.Create)
		storeBargainUserGroup.DELETE("/:storeBargainUserId", storeBargainUserController.Delete)
		storeBargainUserGroup.PUT("/:storeBargainUserId", storeBargainUserController.Put)
		storeBargainUserGroup.GET("/:storeBargainUserId", storeBargainUserController.Get)
		storeBargainUserGroup.PATCH("/:storeBargainUserId", storeBargainUserController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeBargainUserHelpController := controllers.StoreBargainUserHelpController{}
	storeBargainUserHelpGroup := r.Group("/storeBargainUserHelps")
	{
		storeBargainUserHelpGroup.GET("", storeBargainUserHelpController.List)
		storeBargainUserHelpGroup.POST("", storeBargainUserHelpController.Create)
		storeBargainUserHelpGroup.DELETE("/:storeBargainUserHelpId", storeBargainUserHelpController.Delete)
		storeBargainUserHelpGroup.PUT("/:storeBargainUserHelpId", storeBargainUserHelpController.Put)
		storeBargainUserHelpGroup.GET("/:storeBargainUserHelpId", storeBargainUserHelpController.Get)
		storeBargainUserHelpGroup.PATCH("/:storeBargainUserHelpId", storeBargainUserHelpController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCartController := controllers.StoreCartController{}
	storeCartGroup := r.Group("/storeCarts")
	{
		storeCartGroup.GET("", storeCartController.List)
		storeCartGroup.POST("", storeCartController.Create)
		storeCartGroup.DELETE("/:storeCartId", storeCartController.Delete)
		storeCartGroup.PUT("/:storeCartId", storeCartController.Put)
		storeCartGroup.GET("/:storeCartId", storeCartController.Get)
		storeCartGroup.PATCH("/:storeCartId", storeCartController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCategoryController := controllers.StoreCategoryController{}
	storeCategoryGroup := r.Group("/storeCategorys")
	{
		storeCategoryGroup.GET("", storeCategoryController.List)
		storeCategoryGroup.POST("", storeCategoryController.Create)
		storeCategoryGroup.DELETE("/:storeCategoryId", storeCategoryController.Delete)
		storeCategoryGroup.PUT("/:storeCategoryId", storeCategoryController.Put)
		storeCategoryGroup.GET("/:storeCategoryId", storeCategoryController.Get)
		storeCategoryGroup.PATCH("/:storeCategoryId", storeCategoryController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCombinationController := controllers.StoreCombinationController{}
	storeCombinationGroup := r.Group("/storeCombinations")
	{
		storeCombinationGroup.GET("", storeCombinationController.List)
		storeCombinationGroup.POST("", storeCombinationController.Create)
		storeCombinationGroup.DELETE("/:storeCombinationId", storeCombinationController.Delete)
		storeCombinationGroup.PUT("/:storeCombinationId", storeCombinationController.Put)
		storeCombinationGroup.GET("/:storeCombinationId", storeCombinationController.Get)
		storeCombinationGroup.PATCH("/:storeCombinationId", storeCombinationController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCouponController := controllers.StoreCouponController{}
	storeCouponGroup := r.Group("/storeCoupons")
	{
		storeCouponGroup.GET("", storeCouponController.List)
		storeCouponGroup.POST("", storeCouponController.Create)
		storeCouponGroup.DELETE("/:storeCouponId", storeCouponController.Delete)
		storeCouponGroup.PUT("/:storeCouponId", storeCouponController.Put)
		storeCouponGroup.GET("/:storeCouponId", storeCouponController.Get)
		storeCouponGroup.PATCH("/:storeCouponId", storeCouponController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCouponIssueController := controllers.StoreCouponIssueController{}
	storeCouponIssueGroup := r.Group("/storeCouponIssues")
	{
		storeCouponIssueGroup.GET("", storeCouponIssueController.List)
		storeCouponIssueGroup.POST("", storeCouponIssueController.Create)
		storeCouponIssueGroup.DELETE("/:storeCouponIssueId", storeCouponIssueController.Delete)
		storeCouponIssueGroup.PUT("/:storeCouponIssueId", storeCouponIssueController.Put)
		storeCouponIssueGroup.GET("/:storeCouponIssueId", storeCouponIssueController.Get)
		storeCouponIssueGroup.PATCH("/:storeCouponIssueId", storeCouponIssueController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeCouponUserController := controllers.StoreCouponUserController{}
	storeCouponUserGroup := r.Group("/storeCouponUsers")
	{
		storeCouponUserGroup.GET("", storeCouponUserController.List)
		storeCouponUserGroup.POST("", storeCouponUserController.Create)
		storeCouponUserGroup.DELETE("/:storeCouponUserId", storeCouponUserController.Delete)
		storeCouponUserGroup.PUT("/:storeCouponUserId", storeCouponUserController.Put)
		storeCouponUserGroup.GET("/:storeCouponUserId", storeCouponUserController.Get)
		storeCouponUserGroup.PATCH("/:storeCouponUserId", storeCouponUserController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeOrderController := controllers.StoreOrderController{}
	storeOrderGroup := r.Group("/storeOrders")
	{
		storeOrderGroup.GET("", storeOrderController.List)
		storeOrderGroup.POST("", storeOrderController.Create)
		storeOrderGroup.DELETE("/:storeOrderId", storeOrderController.Delete)
		storeOrderGroup.PUT("/:storeOrderId", storeOrderController.Put)
		storeOrderGroup.GET("/:storeOrderId", storeOrderController.Get)
		storeOrderGroup.PATCH("/:storeOrderId", storeOrderController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storePinkController := controllers.StorePinkController{}
	storePinkGroup := r.Group("/storePinks")
	{
		storePinkGroup.GET("", storePinkController.List)
		storePinkGroup.POST("", storePinkController.Create)
		storePinkGroup.DELETE("/:storePinkId", storePinkController.Delete)
		storePinkGroup.PUT("/:storePinkId", storePinkController.Put)
		storePinkGroup.GET("/:storePinkId", storePinkController.Get)
		storePinkGroup.PATCH("/:storePinkId", storePinkController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeProductController := controllers.StoreProductController{}
	storeProductGroup := r.Group("/storeProducts")
	{
		storeProductGroup.GET("", storeProductController.List)
		storeProductGroup.POST("", storeProductController.Create)
		storeProductGroup.DELETE("/:storeProductId", storeProductController.Delete)
		storeProductGroup.PUT("/:storeProductId", storeProductController.Put)
		storeProductGroup.GET("/:storeProductId", storeProductController.Get)
		storeProductGroup.PATCH("/:storeProductId", storeProductController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeProductCateController := controllers.StoreProductCateController{}
	storeProductCateGroup := r.Group("/storeProductCates")
	{
		storeProductCateGroup.GET("", storeProductCateController.List)
		storeProductCateGroup.POST("", storeProductCateController.Create)
		storeProductCateGroup.DELETE("/:storeProductCateId", storeProductCateController.Delete)
		storeProductCateGroup.PUT("/:storeProductCateId", storeProductCateController.Put)
		storeProductCateGroup.GET("/:storeProductCateId", storeProductCateController.Get)
		storeProductCateGroup.PATCH("/:storeProductCateId", storeProductCateController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeProductReplyController := controllers.StoreProductReplyController{}
	storeProductReplyGroup := r.Group("/storeProductReplys")
	{
		storeProductReplyGroup.GET("", storeProductReplyController.List)
		storeProductReplyGroup.POST("", storeProductReplyController.Create)
		storeProductReplyGroup.DELETE("/:storeProductReplyId", storeProductReplyController.Delete)
		storeProductReplyGroup.PUT("/:storeProductReplyId", storeProductReplyController.Put)
		storeProductReplyGroup.GET("/:storeProductReplyId", storeProductReplyController.Get)
		storeProductReplyGroup.PATCH("/:storeProductReplyId", storeProductReplyController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeSeckillController := controllers.StoreSeckillController{}
	storeSeckillGroup := r.Group("/storeSeckills")
	{
		storeSeckillGroup.GET("", storeSeckillController.List)
		storeSeckillGroup.POST("", storeSeckillController.Create)
		storeSeckillGroup.DELETE("/:storeSeckillId", storeSeckillController.Delete)
		storeSeckillGroup.PUT("/:storeSeckillId", storeSeckillController.Put)
		storeSeckillGroup.GET("/:storeSeckillId", storeSeckillController.Get)
		storeSeckillGroup.PATCH("/:storeSeckillId", storeSeckillController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeServiceController := controllers.StoreServiceController{}
	storeServiceGroup := r.Group("/storeServices")
	{
		storeServiceGroup.GET("", storeServiceController.List)
		storeServiceGroup.POST("", storeServiceController.Create)
		storeServiceGroup.DELETE("/:storeServiceId", storeServiceController.Delete)
		storeServiceGroup.PUT("/:storeServiceId", storeServiceController.Put)
		storeServiceGroup.GET("/:storeServiceId", storeServiceController.Get)
		storeServiceGroup.PATCH("/:storeServiceId", storeServiceController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeServiceLogController := controllers.StoreServiceLogController{}
	storeServiceLogGroup := r.Group("/storeServiceLogs")
	{
		storeServiceLogGroup.GET("", storeServiceLogController.List)
		storeServiceLogGroup.POST("", storeServiceLogController.Create)
		storeServiceLogGroup.DELETE("/:storeServiceLogId", storeServiceLogController.Delete)
		storeServiceLogGroup.PUT("/:storeServiceLogId", storeServiceLogController.Put)
		storeServiceLogGroup.GET("/:storeServiceLogId", storeServiceLogController.Get)
		storeServiceLogGroup.PATCH("/:storeServiceLogId", storeServiceLogController.Patch)
	}
	//!!do not delete gen will generate router code at here

	storeVisitController := controllers.StoreVisitController{}
	storeVisitGroup := r.Group("/storeVisits")
	{
		storeVisitGroup.GET("", storeVisitController.List)
		storeVisitGroup.POST("", storeVisitController.Create)
		storeVisitGroup.DELETE("/:storeVisitId", storeVisitController.Delete)
		storeVisitGroup.PUT("/:storeVisitId", storeVisitController.Put)
		storeVisitGroup.GET("/:storeVisitId", storeVisitController.Get)
		storeVisitGroup.PATCH("/:storeVisitId", storeVisitController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemAdminController := controllers.SystemAdminController{}
	systemAdminGroup := r.Group("/systemAdmins")
	{
		systemAdminGroup.GET("", systemAdminController.List)
		systemAdminGroup.POST("", systemAdminController.Create)
		systemAdminGroup.DELETE("/:systemAdminId", systemAdminController.Delete)
		systemAdminGroup.PUT("/:systemAdminId", systemAdminController.Put)
		systemAdminGroup.GET("/:systemAdminId", systemAdminController.Get)
		systemAdminGroup.PATCH("/:systemAdminId", systemAdminController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemAttachmentCategoryController := controllers.SystemAttachmentCategoryController{}
	systemAttachmentCategoryGroup := r.Group("/systemAttachmentCategorys")
	{
		systemAttachmentCategoryGroup.GET("", systemAttachmentCategoryController.List)
		systemAttachmentCategoryGroup.POST("", systemAttachmentCategoryController.Create)
		systemAttachmentCategoryGroup.DELETE("/:systemAttachmentCategoryId", systemAttachmentCategoryController.Delete)
		systemAttachmentCategoryGroup.PUT("/:systemAttachmentCategoryId", systemAttachmentCategoryController.Put)
		systemAttachmentCategoryGroup.GET("/:systemAttachmentCategoryId", systemAttachmentCategoryController.Get)
		systemAttachmentCategoryGroup.PATCH("/:systemAttachmentCategoryId", systemAttachmentCategoryController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemConfigController := controllers.SystemConfigController{}
	systemConfigGroup := r.Group("/systemConfigs")
	{
		systemConfigGroup.GET("", systemConfigController.List)
		systemConfigGroup.POST("", systemConfigController.Create)
		systemConfigGroup.DELETE("/:systemConfigId", systemConfigController.Delete)
		systemConfigGroup.PUT("/:systemConfigId", systemConfigController.Put)
		systemConfigGroup.GET("/:systemConfigId", systemConfigController.Get)
		systemConfigGroup.PATCH("/:systemConfigId", systemConfigController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemConfigTabController := controllers.SystemConfigTabController{}
	systemConfigTabGroup := r.Group("/systemConfigTabs")
	{
		systemConfigTabGroup.GET("", systemConfigTabController.List)
		systemConfigTabGroup.POST("", systemConfigTabController.Create)
		systemConfigTabGroup.DELETE("/:systemConfigTabId", systemConfigTabController.Delete)
		systemConfigTabGroup.PUT("/:systemConfigTabId", systemConfigTabController.Put)
		systemConfigTabGroup.GET("/:systemConfigTabId", systemConfigTabController.Get)
		systemConfigTabGroup.PATCH("/:systemConfigTabId", systemConfigTabController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemFileController := controllers.SystemFileController{}
	systemFileGroup := r.Group("/systemFiles")
	{
		systemFileGroup.GET("", systemFileController.List)
		systemFileGroup.POST("", systemFileController.Create)
		systemFileGroup.DELETE("/:systemFileId", systemFileController.Delete)
		systemFileGroup.PUT("/:systemFileId", systemFileController.Put)
		systemFileGroup.GET("/:systemFileId", systemFileController.Get)
		systemFileGroup.PATCH("/:systemFileId", systemFileController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemGroupController := controllers.SystemGroupController{}
	systemGroupGroup := r.Group("/systemGroups")
	{
		systemGroupGroup.GET("", systemGroupController.List)
		systemGroupGroup.POST("", systemGroupController.Create)
		systemGroupGroup.DELETE("/:systemGroupId", systemGroupController.Delete)
		systemGroupGroup.PUT("/:systemGroupId", systemGroupController.Put)
		systemGroupGroup.GET("/:systemGroupId", systemGroupController.Get)
		systemGroupGroup.PATCH("/:systemGroupId", systemGroupController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemGroupDataController := controllers.SystemGroupDataController{}
	systemGroupDataGroup := r.Group("/systemGroupDatas")
	{
		systemGroupDataGroup.GET("", systemGroupDataController.List)
		systemGroupDataGroup.POST("", systemGroupDataController.Create)
		systemGroupDataGroup.DELETE("/:systemGroupDataId", systemGroupDataController.Delete)
		systemGroupDataGroup.PUT("/:systemGroupDataId", systemGroupDataController.Put)
		systemGroupDataGroup.GET("/:systemGroupDataId", systemGroupDataController.Get)
		systemGroupDataGroup.PATCH("/:systemGroupDataId", systemGroupDataController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemLogController := controllers.SystemLogController{}
	systemLogGroup := r.Group("/systemLogs")
	{
		systemLogGroup.GET("", systemLogController.List)
		systemLogGroup.POST("", systemLogController.Create)
		systemLogGroup.DELETE("/:systemLogId", systemLogController.Delete)
		systemLogGroup.PUT("/:systemLogId", systemLogController.Put)
		systemLogGroup.GET("/:systemLogId", systemLogController.Get)
		systemLogGroup.PATCH("/:systemLogId", systemLogController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemMenusController := controllers.SystemMenusController{}
	systemMenusGroup := r.Group("/systemMenuss")
	{
		systemMenusGroup.GET("", systemMenusController.List)
		systemMenusGroup.POST("", systemMenusController.Create)
		systemMenusGroup.DELETE("/:systemMenusId", systemMenusController.Delete)
		systemMenusGroup.PUT("/:systemMenusId", systemMenusController.Put)
		systemMenusGroup.GET("/:systemMenusId", systemMenusController.Get)
		systemMenusGroup.PATCH("/:systemMenusId", systemMenusController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemNoticeController := controllers.SystemNoticeController{}
	systemNoticeGroup := r.Group("/systemNotices")
	{
		systemNoticeGroup.GET("", systemNoticeController.List)
		systemNoticeGroup.POST("", systemNoticeController.Create)
		systemNoticeGroup.DELETE("/:systemNoticeId", systemNoticeController.Delete)
		systemNoticeGroup.PUT("/:systemNoticeId", systemNoticeController.Put)
		systemNoticeGroup.GET("/:systemNoticeId", systemNoticeController.Get)
		systemNoticeGroup.PATCH("/:systemNoticeId", systemNoticeController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemNoticeAdminController := controllers.SystemNoticeAdminController{}
	systemNoticeAdminGroup := r.Group("/systemNoticeAdmins")
	{
		systemNoticeAdminGroup.GET("", systemNoticeAdminController.List)
		systemNoticeAdminGroup.POST("", systemNoticeAdminController.Create)
		systemNoticeAdminGroup.DELETE("/:systemNoticeAdminId", systemNoticeAdminController.Delete)
		systemNoticeAdminGroup.PUT("/:systemNoticeAdminId", systemNoticeAdminController.Put)
		systemNoticeAdminGroup.GET("/:systemNoticeAdminId", systemNoticeAdminController.Get)
		systemNoticeAdminGroup.PATCH("/:systemNoticeAdminId", systemNoticeAdminController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemRoleController := controllers.SystemRoleController{}
	systemRoleGroup := r.Group("/systemRoles")
	{
		systemRoleGroup.GET("", systemRoleController.List)
		systemRoleGroup.POST("", systemRoleController.Create)
		systemRoleGroup.DELETE("/:systemRoleId", systemRoleController.Delete)
		systemRoleGroup.PUT("/:systemRoleId", systemRoleController.Put)
		systemRoleGroup.GET("/:systemRoleId", systemRoleController.Get)
		systemRoleGroup.PATCH("/:systemRoleId", systemRoleController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemUserLevelController := controllers.SystemUserLevelController{}
	systemUserLevelGroup := r.Group("/systemUserLevels")
	{
		systemUserLevelGroup.GET("", systemUserLevelController.List)
		systemUserLevelGroup.POST("", systemUserLevelController.Create)
		systemUserLevelGroup.DELETE("/:systemUserLevelId", systemUserLevelController.Delete)
		systemUserLevelGroup.PUT("/:systemUserLevelId", systemUserLevelController.Put)
		systemUserLevelGroup.GET("/:systemUserLevelId", systemUserLevelController.Get)
		systemUserLevelGroup.PATCH("/:systemUserLevelId", systemUserLevelController.Patch)
	}
	//!!do not delete gen will generate router code at here

	systemUserTaskController := controllers.SystemUserTaskController{}
	systemUserTaskGroup := r.Group("/systemUserTasks")
	{
		systemUserTaskGroup.GET("", systemUserTaskController.List)
		systemUserTaskGroup.POST("", systemUserTaskController.Create)
		systemUserTaskGroup.DELETE("/:systemUserTaskId", systemUserTaskController.Delete)
		systemUserTaskGroup.PUT("/:systemUserTaskId", systemUserTaskController.Put)
		systemUserTaskGroup.GET("/:systemUserTaskId", systemUserTaskController.Get)
		systemUserTaskGroup.PATCH("/:systemUserTaskId", systemUserTaskController.Patch)
	}
	//!!do not delete gen will generate router code at here

	tokenController := controllers.TokenController{}
	tokenGroup := r.Group("/tokens")
	{
		tokenGroup.GET("", tokenController.List)
		tokenGroup.POST("", tokenController.Create)
		tokenGroup.DELETE("/:tokenId", tokenController.Delete)
		tokenGroup.PUT("/:tokenId", tokenController.Put)
		tokenGroup.GET("/:tokenId", tokenController.Get)
		tokenGroup.PATCH("/:tokenId", tokenController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userController := controllers.UserController{}
	userGroup := r.Group("/users")
	{
		userGroup.GET("", userController.List)
		userGroup.POST("", userController.Create)
		userGroup.DELETE("/:userId", userController.Delete)
		userGroup.PUT("/:userId", userController.Put)
		userGroup.GET("/:userId", userController.Get)
		userGroup.PATCH("/:userId", userController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userAddressController := controllers.UserAddressController{}
	userAddressGroup := r.Group("/userAddresss")
	{
		userAddressGroup.GET("", userAddressController.List)
		userAddressGroup.POST("", userAddressController.Create)
		userAddressGroup.DELETE("/:userAddressId", userAddressController.Delete)
		userAddressGroup.PUT("/:userAddressId", userAddressController.Put)
		userAddressGroup.GET("/:userAddressId", userAddressController.Get)
		userAddressGroup.PATCH("/:userAddressId", userAddressController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userBillController := controllers.UserBillController{}
	userBillGroup := r.Group("/userBills")
	{
		userBillGroup.GET("", userBillController.List)
		userBillGroup.POST("", userBillController.Create)
		userBillGroup.DELETE("/:userBillId", userBillController.Delete)
		userBillGroup.PUT("/:userBillId", userBillController.Put)
		userBillGroup.GET("/:userBillId", userBillController.Get)
		userBillGroup.PATCH("/:userBillId", userBillController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userEnterController := controllers.UserEnterController{}
	userEnterGroup := r.Group("/userEnters")
	{
		userEnterGroup.GET("", userEnterController.List)
		userEnterGroup.POST("", userEnterController.Create)
		userEnterGroup.DELETE("/:userEnterId", userEnterController.Delete)
		userEnterGroup.PUT("/:userEnterId", userEnterController.Put)
		userEnterGroup.GET("/:userEnterId", userEnterController.Get)
		userEnterGroup.PATCH("/:userEnterId", userEnterController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userExtractController := controllers.UserExtractController{}
	userExtractGroup := r.Group("/userExtracts")
	{
		userExtractGroup.GET("", userExtractController.List)
		userExtractGroup.POST("", userExtractController.Create)
		userExtractGroup.DELETE("/:userExtractId", userExtractController.Delete)
		userExtractGroup.PUT("/:userExtractId", userExtractController.Put)
		userExtractGroup.GET("/:userExtractId", userExtractController.Get)
		userExtractGroup.PATCH("/:userExtractId", userExtractController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userGroupController := controllers.UserGroupController{}
	userGroupGroup := r.Group("/userGroups")
	{
		userGroupGroup.GET("", userGroupController.List)
		userGroupGroup.POST("", userGroupController.Create)
		userGroupGroup.DELETE("/:userGroupId", userGroupController.Delete)
		userGroupGroup.PUT("/:userGroupId", userGroupController.Put)
		userGroupGroup.GET("/:userGroupId", userGroupController.Get)
		userGroupGroup.PATCH("/:userGroupId", userGroupController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userLevelController := controllers.UserLevelController{}
	userLevelGroup := r.Group("/userLevels")
	{
		userLevelGroup.GET("", userLevelController.List)
		userLevelGroup.POST("", userLevelController.Create)
		userLevelGroup.DELETE("/:userLevelId", userLevelController.Delete)
		userLevelGroup.PUT("/:userLevelId", userLevelController.Put)
		userLevelGroup.GET("/:userLevelId", userLevelController.Get)
		userLevelGroup.PATCH("/:userLevelId", userLevelController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userNoticeController := controllers.UserNoticeController{}
	userNoticeGroup := r.Group("/userNotices")
	{
		userNoticeGroup.GET("", userNoticeController.List)
		userNoticeGroup.POST("", userNoticeController.Create)
		userNoticeGroup.DELETE("/:userNoticeId", userNoticeController.Delete)
		userNoticeGroup.PUT("/:userNoticeId", userNoticeController.Put)
		userNoticeGroup.GET("/:userNoticeId", userNoticeController.Get)
		userNoticeGroup.PATCH("/:userNoticeId", userNoticeController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userNoticeSeeController := controllers.UserNoticeSeeController{}
	userNoticeSeeGroup := r.Group("/userNoticeSees")
	{
		userNoticeSeeGroup.GET("", userNoticeSeeController.List)
		userNoticeSeeGroup.POST("", userNoticeSeeController.Create)
		userNoticeSeeGroup.DELETE("/:userNoticeSeeId", userNoticeSeeController.Delete)
		userNoticeSeeGroup.PUT("/:userNoticeSeeId", userNoticeSeeController.Put)
		userNoticeSeeGroup.GET("/:userNoticeSeeId", userNoticeSeeController.Get)
		userNoticeSeeGroup.PATCH("/:userNoticeSeeId", userNoticeSeeController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userRechargeController := controllers.UserRechargeController{}
	userRechargeGroup := r.Group("/userRecharges")
	{
		userRechargeGroup.GET("", userRechargeController.List)
		userRechargeGroup.POST("", userRechargeController.Create)
		userRechargeGroup.DELETE("/:userRechargeId", userRechargeController.Delete)
		userRechargeGroup.PUT("/:userRechargeId", userRechargeController.Put)
		userRechargeGroup.GET("/:userRechargeId", userRechargeController.Get)
		userRechargeGroup.PATCH("/:userRechargeId", userRechargeController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userSignController := controllers.UserSignController{}
	userSignGroup := r.Group("/userSigns")
	{
		userSignGroup.GET("", userSignController.List)
		userSignGroup.POST("", userSignController.Create)
		userSignGroup.DELETE("/:userSignId", userSignController.Delete)
		userSignGroup.PUT("/:userSignId", userSignController.Put)
		userSignGroup.GET("/:userSignId", userSignController.Get)
		userSignGroup.PATCH("/:userSignId", userSignController.Patch)
	}
	//!!do not delete gen will generate router code at here

	userTaskFinishController := controllers.UserTaskFinishController{}
	userTaskFinishGroup := r.Group("/userTaskFinishs")
	{
		userTaskFinishGroup.GET("", userTaskFinishController.List)
		userTaskFinishGroup.POST("", userTaskFinishController.Create)
		userTaskFinishGroup.DELETE("/:userTaskFinishId", userTaskFinishController.Delete)
		userTaskFinishGroup.PUT("/:userTaskFinishId", userTaskFinishController.Put)
		userTaskFinishGroup.GET("/:userTaskFinishId", userTaskFinishController.Get)
		userTaskFinishGroup.PATCH("/:userTaskFinishId", userTaskFinishController.Patch)
	}
	//!!do not delete gen will generate router code at here

	wechatMessageController := controllers.WechatMessageController{}
	wechatMessageGroup := r.Group("/wechatMessages")
	{
		wechatMessageGroup.GET("", wechatMessageController.List)
		wechatMessageGroup.POST("", wechatMessageController.Create)
		wechatMessageGroup.DELETE("/:wechatMessageId", wechatMessageController.Delete)
		wechatMessageGroup.PUT("/:wechatMessageId", wechatMessageController.Put)
		wechatMessageGroup.GET("/:wechatMessageId", wechatMessageController.Get)
		wechatMessageGroup.PATCH("/:wechatMessageId", wechatMessageController.Patch)
	}
	//!!do not delete gen will generate router code at here

}
