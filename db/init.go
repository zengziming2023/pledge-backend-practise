package db

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var MySql *gorm.DB
var RedisPool *redis.Pool
