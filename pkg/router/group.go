package router

import (
	"github.com/gin-gonic/gin"
	"study/testdeploy/pkg/handlers"
)

// 客户端分组
func clientGroup(engine *gin.Engine) {
	router := engine.Group("/v1/client/gameconfigapisrv")
	{
		{
			router.GET("/gameconfig/getConfig", handlers.OnGetConfigData)
		}
	}
}

// 运营平台分组
func omsGroup(engine *gin.Engine) {
	router := engine.Group("/v1/oms/gameconfigapisrv")
	{
		// 游戏配置管理接口
		{
			router.POST("/gameconfig/add", handlers.OnAddGameConfig)
			router.POST("/gameconfig/delete", handlers.OnDeleteGameConfig)
			router.POST("/gameconfig/update", handlers.OnUpdateGameConfig)
			router.GET("/gameconfig/list", handlers.OnListGameConfig)
		}
	}
}
