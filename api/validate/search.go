package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
)

type Search struct{}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Search(ctx *gin.Context, req *request.Search) int {
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

	if req.ChainID != 97 && req.ChainID != 56 {
		return statuscode.ChainIdErr
	}

	return statuscode.CommonSuccess
}
