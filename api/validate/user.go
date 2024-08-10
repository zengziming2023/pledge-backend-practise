package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) Login(ctx *gin.Context, req *request.Login) int {
	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return statuscode.ParameterEmptyErr
	} else if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			if e.Field() == "name" && e.Tag() == "required" {
				return statuscode.PNameEmpty
			}
			if e.Field() == "password" && e.Tag() == "required" {
				return statuscode.PNameEmpty
			}
		}
		return statuscode.CommonErrServerErr
	}

	return statuscode.CommonSuccess
}
