package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"study/testdeploy/pkg/config"
	"time"
)

var (
	ConfigRedisPool *redis.Pool
)

func NewVipRedisConn(cacheName string, tomlConfig *config.Config) gin.HandlerFunc {
	cacheConfig, ok := tomlConfig.RedisServerConf(cacheName)
	if !ok {
		panic(fmt.Sprintf("%v not set.", cacheName))
	}

	// 链接数据库
	ConfigRedisPool = newPool(cacheConfig.Addr, cacheConfig.Password, cacheConfig.DB)

	return func(c *gin.Context) {
		c.Set(cacheName, ConfigRedisPool)
		c.Next()
	}

}

// newPool New redis pool.
func newPool(server, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// getExpireAt 获取过期时间
func getExpireAt(day int) (deadLine int32) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	end := start.AddDate(0, 0, day)

	deadLine = int32(end.Unix())
	return
}
