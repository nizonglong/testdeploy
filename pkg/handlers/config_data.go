package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"study/testdeploy/pkg/define"
	"study/testdeploy/pkg/models/psql"
)

func OnGetConfigData(c *gin.Context) {
	strGameID := c.Query(define.StrGameID)
	intGameID, err := strconv.Atoi(strGameID)
	strType := c.Query(define.StrType)
	intType, err := strconv.Atoi(strType)
	// param judge
	if err != nil {
		//logging.Errorf("OnGetConfigData err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.ErrParam, nil))
		return
	}
	// 获取json字符串data
	gameConfigs, err := psql.ListGameConfig(psql.QiPaiDB, int16(intGameID), int16(intType))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			//logging.Errorf("OnGetConfigData->ListGameConfig->empty!")
			c.JSON(http.StatusOK, result(define.IntOne, define.StatusOk, nil))
			return
		}
		//logging.Errorf("OnGetConfigData->ListGameConfig err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.RetryLaterError, nil))
		return
	}

	data, err := translateConfigData(gameConfigs, intType)
	if err != nil {
		//logging.Debugf("OnGetConfigData->translateConfigData err:%v", err)
		c.JSON(http.StatusOK, result(define.IntOne, define.RetryLaterError, data))
		return
	}

	//logging.Debugf("OnGetConfigData success")
	c.JSON(http.StatusOK, result(define.IntZero, define.StatusOk, data))
	return
}

func translateConfigData(gameConfigs []*define.GameConfig, configType int) (data []*interface{}, err error) {
	for _, config := range gameConfigs {
		temp := getData(configType)
		err = json.Unmarshal([]byte(config.Data), &temp)
		if err != nil {
			//logging.Errorf("translateConfigData->Unmarshal err:%v", err)
			return nil, err
		}
		data = append(data, &temp)
	}

	return data, nil
}

// 对配置进行类型转换m
func getData(configType int) (data interface{}) {
	switch configType {
	case define.ConfigTypeLoading:
		data = &define.LoadingWordConfig{}
		break
	case define.ConfigTypeMarquee:
		data = &define.MarqueeConfig{}
		break
	case define.ConfigTypeADCoin:
		data = &[]define.ADWatchCoinConfig{}
		break
	case define.ConfigTypeShopBanner:
		data = &define.BannerConfig{}
		break
	case define.ConfigTypeEmail:
		data = &define.EmailConfig{}
		break
	case define.ConfigTypeCommunity:
		data = &define.CommunityConfig{}
		break
	case define.ConfigTypeShopDailyFreeCoin:
		break
	case define.ConfigTypeADCount:
		data = &[]define.ADWatchCountConfig{}
		break
	case define.ConfigTypeTimeZone:
		data = &define.TimeZone{}
		break
	default:
		return nil
	}
	return data
}
