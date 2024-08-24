package main

import (
	"pledge-backend-practise/db"
	"pledge-backend-practise/db/tasks"
)

func main() {
	// init mysql
	db.InitMySql()
	// init redis
	db.InitRedis()
	// pool task
	tasks.StartTask()
}
