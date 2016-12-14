package cache

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/cache/redis"
	"time"
	"github.com/garyburd/redigo/redis"
)

var (
	// 定义常量
	RedisClient     *redis.Pool
)


// Put put cache to redis.
func Put(key string, val interface{}, timeout time.Duration) error {
	var err error

	if _, err = RedisClient.Get().Do("SET", key, val); err != nil {
		return err
	}
	if _, err = RedisClient.Get().Do("SETEX", key, int64(timeout / time.Second), val); err != nil {
		return err
	}
	return err
}

// Get cache from redis.
func Get(key string) interface{} {
	if v, err := RedisClient.Get().Do("GET", key); err == nil {
		return v
	}
	return nil
}

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := "127.0.0.1:9706"
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST, redis.DialPassword("1qaz2wsx3edc"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},

	}
}
