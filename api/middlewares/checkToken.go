package middlewares

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/config"
	"pledge-backend-practise/db"
	"pledge-backend-practise/utils"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authCode")

		userName, err := utils.ParseToken(token, config.Config.Jwt.SecretKey)
		if err != nil {
			c.AbortWithStatus(statuscode.TokenErr)
			return
		}

		if userName != config.Config.DefaultAdmin.Username {
			c.AbortWithStatus(statuscode.TokenErr)
			return
		}

		// Judge whether the user logout
		loginStatus, err := db.RedisGetString(userName)
		if loginStatus != "login_ok" {
			response.JsonResponse(c, statuscode.TokenErr, nil)
			c.Abort()
			return
		}

		c.Set("userName", userName)
		c.Next()
	}
}
