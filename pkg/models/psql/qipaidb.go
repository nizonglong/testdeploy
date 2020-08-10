package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"study/testdeploy/pkg/config"
	"study/testdeploy/pkg/define"
)

// 创建qipaidb连接
func NewQipaiDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	QiPaiDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	QiPaiDB.DB().SetMaxOpenConns(10)
}

/**
 * -------------------------------------------------
 * 游戏配置SQL
 * -------------------------------------------------
 */
func AddGameConfig(qipaidb *gorm.DB, gameConfig *define.GameConfig) (err error) {
	err = qipaidb.Debug().Table(define.TableGameConfig).Create(gameConfig).Error
	return
}

func DeleteGameConfig(qipaidb *gorm.DB, gameID, gameConfigID int16) (err error) {
	gameConfig := &define.GameConfig{ID: gameConfigID}
	err = qipaidb.Debug().Table(define.TableGameConfig).Where("game_id = ?", gameID).
		Delete(gameConfig).Error
	return
}

func UpdateGameConfig(qipaidb *gorm.DB, gameConfig *define.GameConfig) (err error) {
	err = qipaidb.Debug().Table(define.TableGameConfig).Save(gameConfig).Error
	return
}

func ListGameConfig(qipaidb *gorm.DB, gameID, typez int16) (gameConfigs []*define.GameConfig, err error) {
	err = qipaidb.Debug().Table(define.TableGameConfig).Where("game_id = ? and type = ?", gameID, typez).Order("id").
		Find(&gameConfigs).Error
	return
}
