package controllers

import (
	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

func (crtl *StoreCategoryController) PidByCategory(c *gin.Context) {
	storeCategory := &models.StoreCategory{}
	list, err := storeCategory.PidByCategory(0)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, list)
}

func (crtl *StoreCategoryController) GetProductCategory(c *gin.Context) {
	storeCategory := &models.StoreCategory{}
	rdata, err := storeCategory.GetProductCategory()
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, rdata)
}
