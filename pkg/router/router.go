package router

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"runtime"
	"study/testdeploy/pkg/config"
	"study/testdeploy/pkg/handlers"
	"study/testdeploy/pkg/middlewares"
	"time"
)

var (
	tomlFilePath, mode string
)

// InitEngine 初始化engine
func InitEngine() (engine *gin.Engine, tomlConfig *config.Config, err error) {
	flag.StringVar(&tomlFilePath, "config", "docs/local.toml", "服务配置文件")
	flag.StringVar(&mode, "mode", "release", "模型-debug还是release还是test")
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	gin.SetMode(mode)

	// 解析配置文件
	config.Conf, err = config.UnmarshalConfig(tomlFilePath)
	if err != nil {
		//logging.Errorf("UnmarshalConfig: err:%v\n", err)
	}

	// 绑定路由，及公共的tomlConfig
	engine = gin.New()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"x-xq5-jwt", "Download-Status"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(gin.Recovery())

	// 加载中间件
	userMiddleware(engine)

	// 路由配置
	engine.GET("/", handlers.Index) // 首页

	// 运营平台分组
	omsGroup(engine)

	// 客户端分组
	clientGroup(engine)

	return engine, config.Conf, nil
}

// userMiddleware 设置中间件
func userMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.Config(config.Conf))

}
