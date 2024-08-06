package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/services"
	"pledge-backend-practise/api/validate"
)

type PoolController struct{}

func (c *PoolController) PoolBaseInfo(ctx *gin.Context) {
	req := request.PoolBaseInfo{}
	var result []models.PoolBaseInfoRes

	// 参数校验
	errorCode := validate.NewPoolBaseInfo().PoolBaseInfo(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		ctx.JSON(errorCode, nil)
		return
	}

	errorCode = services.NewPool().PoolBaseInfo(req.ChainId, &result)
	if errorCode != statuscode.CommonSuccess {
		ctx.JSON(errorCode, nil)
		return
	}

	ctx.JSON(statuscode.CommonSuccess, result)

}
