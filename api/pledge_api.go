package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pledge-backend-practise/api/middlewares"
	"pledge-backend-practise/api/models/kucoin"
	"pledge-backend-practise/api/models/ws"
	"pledge-backend-practise/api/routes"
	"pledge-backend-practise/api/static"
	"pledge-backend-practise/api/validate"
	"pledge-backend-practise/config"
	"pledge-backend-practise/db"
	"pledge-backend-practise/log"
)

func main() {
	log.Logger.Info("pledge api")

	//init mysql
	db.InitMySql()
	//init redis
	db.InitRedis()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// websocket server
	go ws.StartServer()

	// get plgr price from kucoin-exchange
	go kucoin.GetExchangePrice()

	// gin start
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	defer func(app *gin.Engine, port string) {
		err := app.Run(port)
		if err != nil {
			log.Logger.Error("app run error", zap.Error(err))
		}
	}(app, ":"+config.Config.Env.Port)
	log.Logger.Info(config.Config.Env.Port)

	staticImgPath := static.GetCurrentAbPathByCaller()
	log.Logger.Info("static img path: " + staticImgPath)
	app.Static("/storage/", staticImgPath)
	app.Use(middlewares.Cores()) // 「 Cross domain Middleware 」
	routes.InitRouter(app)
}
