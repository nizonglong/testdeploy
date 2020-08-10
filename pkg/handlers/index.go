package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// api交互通用返回结构
type RespInfo struct {
	ErrNo  int8        `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// 返回赋值
func result(code int8, msg string, data interface{}) *RespInfo {
	return &RespInfo{
		ErrNo:  code,
		ErrMsg: msg,
		Data:   data,
	}
}

// Index 首页
func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
