package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
)

type TokenList struct{}

func NewTokenList() *TokenList {
	return &TokenList{}
}

func (t *TokenList) TokenList(ctx *gin.Context, req *request.TokenList) int {
	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return statuscode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "chainId" && e.Tag() == "required" {
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
