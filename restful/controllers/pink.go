package controllers

import (
	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

type PinkController struct {
}

//获取拼团列表
func (ctl *PinkController) CombinationList(c *gin.Context) {

	storeCombination := &models.StoreCombination{}

	req := &models.Query{
		Page:    1,
		PageNum: 20,
	}

	list, err := storeCombination.GetAll(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	handleOk(c, list)

}
