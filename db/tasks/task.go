package tasks

import (
	"pledge-backend-practise/api/services"
	"pledge-backend-practise/db"
	"pledge-backend-practise/db/common"
	"pledge-backend-practise/log"
)

func StartTask() {
	log.Logger.Info("Starting task: " + common.PlgrAdminPrivateKey)

	// clean redis first.
	if err := db.RedisFlushDB(); err != nil {
		log.Logger.Error("Error while flushing redis for task: " + err.Error())
		return
	}

	//init task
	services.NewPool().UpdateAllPoolInfo()
}
