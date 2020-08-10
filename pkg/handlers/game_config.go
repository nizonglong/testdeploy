package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"study/testdeploy/pkg/define"
	"study/testdeploy/pkg/models/psql"
)

// 添加配置
func OnAddGameConfig(c *gin.Context) {
	req := &define.GameConfig{}
	if err := c.Bind(req); err != nil {
		//logging.Errorf("OnAddGameConfig: err: %v", err.Error())
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrParam, nil))
		return
	}

	if err := psql.AddGameConfig(psql.QiPaiDB, req); err != nil {
		//logging.Errorf("OnAddGameConfig->AddGameConfig err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrDB, nil))
		return
	}

	c.JSON(http.StatusOK, result(define.IntZero, define.StatusOk, nil))
	return
}

// 删除配置
func OnDeleteGameConfig(c *gin.Context) {
	req := &define.DeleteGameConfigRequest{}
	if err := c.Bind(req); err != nil {
		//logging.Errorf("OnDeleteGameConfig: err: %v", err.Error())
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrParam, nil))
		return
	}

	if req.GameConfigID <= 0 {
		//logging.Errorf("OnDeleteGameConfig->DeleteGameConfig GameConfigID 不能为孔")
		c.JSON(http.StatusOK, result(define.IntOne, define.DataNotFit, nil))
		return
	}

	if err := psql.DeleteGameConfig(psql.QiPaiDB, req.GameID, req.GameConfigID); err != nil {
		//logging.Errorf("OnDeleteGameConfig->DeleteGameConfig err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrDB, nil))
		return
	}

	c.JSON(http.StatusOK, result(define.IntZero, define.StatusOk, nil))
	return
}

// 更新配置
func OnUpdateGameConfig(c *gin.Context) {
	req := &define.GameConfig{}
	if err := c.Bind(req); err != nil {
		//logging.Errorf("OnUpdateGameConfig: err: %v", err.Error())
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrParam, nil))
		return
	}

	if err := psql.UpdateGameConfig(psql.QiPaiDB, req); err != nil {
		//logging.Errorf("OnUpdateGameConfig->UpdateGameConfig err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrDB, nil))
		return
	}

	c.JSON(http.StatusOK, result(define.IntZero, define.StatusOk, nil))
	return
}

// 获取配置
func OnListGameConfig(c *gin.Context) {
	strGameID := c.Query(define.StrGameID)
	intGameID, err := strconv.Atoi(strGameID)
	int16GameID := int16(intGameID)

	strType := c.Query(define.StrType)
	intType, err := strconv.Atoi(strType)
	int16Type := int16(intType)
	if err != nil {
		//logging.Errorf("OnListGameConfig-> 参数转换 int16 失败, err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, "参数转换出错", nil))
		return
	}

	gameConfigs, err := psql.ListGameConfig(psql.QiPaiDB, int16GameID, int16Type)
	if err != nil {
		//logging.Errorf("OnListGameConfig->ListGameConfig err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrDB, nil))
		return
	}

	c.JSON(http.StatusOK, result(define.IntZero, define.StatusOk, gameConfigs))
	return
}
