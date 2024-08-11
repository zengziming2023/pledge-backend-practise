package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
)

type MultiSign struct{}

func NewMultiSign() *MultiSign {
	return &MultiSign{}
}

func (m *MultiSign) SetMultiSign(ctx *gin.Context, req *request.SetMultiSign) int {

	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return statuscode.ParameterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, err := range errors {
			if err.Field() == "SpName" && err.Tag() == "required" {
				return statuscode.PNameEmpty
			}
		}
		return statuscode.CommonErrServerErr
	}
	if req.ChainId != 97 && req.ChainId != 56 {
		return statuscode.ChainIdErr
	}

	return statuscode.CommonSuccess
}

func (m *MultiSign) GetMultiSign(ctx *gin.Context, req *request.GetMultiSign) int {
	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return statuscode.ParameterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, err := range errors {
			if err.Field() == "ChainId" && err.Tag() == "required" {
				return statuscode.ChainIdEmpty
			}
		}
		return statuscode.CommonErrServerErr
	}

	if req.ChainId != 97 && req.ChainId != 56 {
		return statuscode.ChainIdErr
	}

	return statuscode.CommonSuccess
}
