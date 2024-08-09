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
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	// 业务处理
	errorCode = services.NewPool().PoolBaseInfo(req.ChainId, &result)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, result)

}

func (c *PoolController) PoolDataInfo(ctx *gin.Context) {
	req := request.PoolDataInfo{}
	var result *[]response.PoolDataRes

	// 参数校验
	errorCode := validate.NewPoolDataInfo().PoolDataInfo(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	// 业务处理 service 取数据
	errorCode = services.NewPool().PoolDataInfo(req.ChainId, result)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, result)
}

func (c *PoolController) TokenList(ctx *gin.Context) {
	req := request.TokenList{}
	var rsp *response.TokenList

	// 参数校验
	errorCode := validate.NewTokenList().TokenList(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	errorCode = services.NewTokenList().TokenList(&req, rsp)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, rsp)
}
