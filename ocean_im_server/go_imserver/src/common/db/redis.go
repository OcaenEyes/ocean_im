package db

import (
	"github.com/gomodule/redigo/redis"
	"oceanim/go_server/src/config"
	"time"
)

var RedisPool *redis.Pool

func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.Conf.Redis.DBMaxIdle,
		MaxActive:   config.Conf.Redis.DBMaxActive,
		IdleTimeout: time.Duration(config.Conf.Redis.DBIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				config.Conf.Redis.DBAddress,
				redis.DialConnectTimeout(time.Duration(1000) * time.Millisecond),
				redis.DialReadTimeout(time.Duration(1000) * time.Millisecond),
				redis.DialDatabase(0),
				redis.DialPassword(config.Conf.Redis.DBPassWord))
		},
	}

}
