package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/api/services"
	"pledge-backend-practise/api/validate"
)

type PoolController struct{}

func (c *PoolController) PoolBaseInfo(ctx *gin.Context) {
	req := request.PoolBaseInfo{}
	var result []response.PoolBaseInfoRes

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

func (c *PoolController) PoolDataInfo(ctx *gin.Context) {
	req := request.PoolDataInfo{}
	var result *[]response.PoolDataRes

	// 参数校验
	errorCode := validate.NewPoolDataInfo().PoolDataInfo(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		ctx.JSON(errorCode, nil)
		return
	}

	// service 取数据
	errorCode = services.NewPool().PoolDataInfo(req.ChainId, result)
	if errorCode != statuscode.CommonSuccess {
		ctx.JSON(errorCode, nil)
		return
	}

	ctx.JSON(statuscode.CommonSuccess, result)
}
