package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/api/services"
	"pledge-backend-practise/api/validate"
)

type multiSignPoolController struct{}

func NewMultiSignPoolController() *multiSignPoolController {
	return &multiSignPoolController{}
}

func (c *multiSignPoolController) SetMultiSign(ctx *gin.Context) {
	req := request.SetMultiSign{}

	errorCode := validate.NewMultiSign().SetMultiSign(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	errorCode = services.NewMultiSign().SetMultiSign(&req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, nil)
}

func (c *multiSignPoolController) GetMultiSign(ctx *gin.Context) {
	req := request.GetMultiSign{}
	rsp := response.MultiSign{}

	errorCode := validate.NewMultiSign().GetMultiSign(ctx, &req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	errorCode = services.NewMultiSign().GetMultiSign(&req, &rsp)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, rsp)
}
