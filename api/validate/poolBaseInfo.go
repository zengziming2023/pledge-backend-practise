package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
)

type PoolBaseInfo struct{}

func NewPoolBaseInfo() *PoolBaseInfo {
	return &PoolBaseInfo{}
}

func (info *PoolBaseInfo) PoolBaseInfo(c *gin.Context, req *request.PoolBaseInfo) int {
	err := c.ShouldBind(req)
	if err == io.EOF {
		return statuscode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
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
