package middlewares

import (
	"github.com/gin-gonic/gin"
	"study/testdeploy/pkg/config"
	"study/testdeploy/pkg/define"
)

func Config(tomlConfig *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(define.StrConfig, tomlConfig)
		c.Next()
	}
}
