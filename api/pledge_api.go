package main

import (
	"pledge-backend-practise/config"
	"pledge-backend-practise/log"
)

func main() {
	log.Logger.Info("pledge api")

	log.Logger.Info(config.Config.Env.Port)
}
