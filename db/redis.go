package db

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pledge-backend-practise/config"
	"pledge-backend-practise/log"
	"time"
)

func InitRedis() *redis.Pool {
	log.Logger.Info("init redis.")
	redisConf := config.Config.Redis
	// 建立连接池
	RedisPool = &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,   // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   redisConf.MaxActive, // 最大的激活连接数，表示同时最多有N个连接   0 表示无穷大
		Wait:        true,                // 如果连接数不足则阻塞等待
		IdleTimeout: time.Duration(redisConf.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConf.Address, redisConf.Port))
			if err != nil {
				return nil, err
			}
			// 验证密码
			_, err = conn.Do("auth", redisConf.Password)
			if err != nil {
				panic("redis auth failed: " + err.Error())
			}
			// 选择db
			_, err = conn.Do("select", redisConf.Db)
			if err != nil {
				panic("redis select db err " + err.Error())
			}

			return conn, nil
		},
	}

	err := RedisPool.Get().Err()
	if err != nil {
		//panic("redis connect failed: " + err.Error())
		log.Logger.Error("init redis pool error: " + err.Error())
		return nil
	}

	return RedisPool
}

// RedisSet 设置key、value、time
func RedisSet(key string, value interface{}, aliveSeconds int) error {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var err error
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if aliveSeconds > 0 {
		_, err = conn.Do("SET", key, valueBytes, "EX", aliveSeconds)

	} else {
		_, err = conn.Do("SET", key, valueBytes)

	}
	if err != nil {
		return err
	}
	return nil
}

// RedisSetString  设置key、value、time
func RedisSetString(key string, value string, aliveSeconds int) error {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var err error
	if aliveSeconds > 0 {
		_, err = redis.String(conn.Do("SET", key, value, "EX", aliveSeconds))
	} else {
		_, err = redis.String(conn.Do("SET", key, value))
	}

	if err != nil {
		return err
	}
	return nil
}

func RedisGet(key string) ([]byte, error) {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	bytes, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func RedisGetString(key string) (string, error) {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()
	str, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return str, nil
}

func RedisSetInt64(key string, value int64, aliveSeconds int) error {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var err error
	if aliveSeconds > 0 {
		_, err = redis.Int64(conn.Do("SET", key, value, "EX", aliveSeconds))
	} else {
		_, err = redis.Int64(conn.Do("SET", key, value))
	}
	if err != nil {
		return err
	}

	return nil
}

func RedisGetInt64(key string) (int64, error) {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	result, err := redis.Int64(conn.Do("GET", key))
	if err != nil {
		return -1, err
	}
	return result, nil
}

func RedisDelete(key string) error {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

// RedisFlushDB 清空当前DB
func RedisFlushDB() error {
	conn := RedisPool.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("FLUSHDB")
	if err != nil {
		return err
	}
	return nil
}
