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
	v2Group.GET("/poolDataInfo", poolController.PoolDataInfo)
	v2Group.GET("/token", poolController.TokenList)
	v2Group.POST("/pool/debtTokenList", middlewares.CheckToken(), poolController.DebtTokenList)
	v2Group.POST("/pool/search", middlewares.CheckToken(), poolController.Search)

	priceController := controllers.NewPriceController()
	// plgr-usdt price
	v2Group.GET("/price", priceController.NewPrice)

	// pledge-defi admin backend
	v2Group.POST("/pool/setMultiSign", middlewares.CheckToken(), nil)
	v2Group.POST("/pool/getMultiSign", middlewares.CheckToken(), nil)

	userController := controllers.NewUserController()
	// user
	v2Group.POST("/user/login", userController.Login)
	v2Group.POST("/user/logout", middlewares.CheckToken(), nil)

}
