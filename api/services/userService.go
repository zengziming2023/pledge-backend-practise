package services

import (
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/config"
	"pledge-backend-practise/db"
	"pledge-backend-practise/log"
	"pledge-backend-practise/utils"
)

type UserService struct {
}

func NewUser() *UserService {
	return &UserService{}
}

func (us *UserService) Login(req *request.Login, rsp *response.Login) int {

	if req.Name == "admin" && req.Password == "password" {
		token, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Error("CreateToken: " + err.Error())
			return statuscode.CommonErrServerErr
		}
		rsp.TokenId = token

		// save to redis
		_ = db.RedisSet(req.Name, "login_ok", config.Config.Jwt.ExpireTime)

		return statuscode.CommonSuccess
	} else {
		return statuscode.NameOrPasswordErr
	}
}

func (us *UserService) Logout(name string) int {
	// remove from redis
	_ = db.RedisDelete(name)

	return statuscode.CommonSuccess
}
