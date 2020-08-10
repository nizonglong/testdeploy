package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	Conf *Config
)

// Config 对应配置文件结构
type Config struct {
	Listen                string                     `toml:"listen"`
	DBServers             map[string]DBServer    `toml:"dbservers"`
	RedisServers          map[string]RedisServer `toml:"redisservers"`
}

func NewConfig(conf *Config) *Config {
	Conf = conf
	return Conf
}

//RedisServer 表示 redis 服务器配置
type RedisServer struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"middlewares"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	if _, err := toml.DecodeFile(tomlfile, &Conf); err != nil {
		return Conf, err
	}
	return Conf, nil
}

// GetListenAddr 监听地址
func (c Config) GetListenAddr() string {
	return c.Listen
}

// RedisServerConf 获取数据库配置
func (c Config) RedisServerConf(key string) (RedisServer, bool) {
	s, ok := c.RedisServers[key]
	return s, ok
}

// NewPool 初始化redis连接
func (c RedisServer) NewPool(maxIdle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", c.Addr, redis.DialDatabase(c.DB), redis.DialPassword(c.Password))
			if err != nil {
				return nil,err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}


// DBServer 表示DB服务器配置
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// ConnectString 表示连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host, m.Port, m.User, m.Password, m.DBName)
}

// DBServerConf 获取数据库配置
func (c Config) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

// NewGormDB 初始化gormdb连接
func (m DBServer) NewGormDB(openConnection int) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", m.ConnectString())
	if err != nil {
		return
	}
	// 设置最大链接数
	db.DB().SetMaxOpenConns(openConnection)
	return
}
