package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pledge-backend-practise/config"
	"pledge-backend-practise/log"
)

func main() {
	log.Logger.Info("pledge api")

	//init mysql

	//init redis

	//gin bind go-playground-validator

	// websocket server

	// get plgr price from kucoin-exchange

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

	//app.Static("/storage/", staticPath)
	//app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	//routes.InitRoute(app)
	app.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{})
	})

}
