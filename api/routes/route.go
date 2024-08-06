package routes

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-practise/api/controllers"
	"pledge-backend-practise/api/middlewares"
	"pledge-backend-practise/config"
)

func InitRouter(e *gin.Engine) {
	v2Group := e.Group("/api/v" + config.Config.Env.Version)

	poolController := &controllers.PoolController{}
	// pledge-defi backend
	v2Group.GET("/poolBaseInfo", poolController.PoolBaseInfo)
	v2Group.GET("/poolDataInfo", nil)
	v2Group.GET("/token", nil)
	v2Group.POST("/pool/debtTokenList", middlewares.CheckToken(), nil)
	v2Group.POST("/pool/search", middlewares.CheckToken(), nil)

	// plgr-usdt price
	v2Group.GET("/price", nil)

	// pledge-defi admin backend
	v2Group.POST("/pool/setMultiSign", middlewares.CheckToken(), nil)
	v2Group.POST("/pool/getMultiSign", middlewares.CheckToken(), nil)

	// user
	v2Group.POST("/user/login", nil)
	v2Group.POST("/user/logout", middlewares.CheckToken(), nil)

}
