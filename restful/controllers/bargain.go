package controllers

import (
	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

type BargainController struct {
}

func (ctl *BargainController) GetBargainList(c *gin.Context) {

	req := new(models.Query)
	err := c.ShouldBindQuery(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeBargain := &models.StoreBargain{}
	list, err := storeBargain.GetList(req)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, list)
}
