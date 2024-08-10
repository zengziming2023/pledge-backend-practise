package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/api/services"
	"pledge-backend-practise/api/validate"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Login(ctx *gin.Context) {
	req := &request.Login{}
	rsp := &response.Login{}

	// 参数校验
	errorCode := validate.NewUser().Login(ctx, req)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	// 业务处理
	errorCode = services.NewUser().Login(req, rsp)
	if errorCode != statuscode.CommonSuccess {
		response.JsonResponse(ctx, errorCode, nil)
		return
	}

	response.JsonResponse(ctx, statuscode.CommonSuccess, rsp)
}
