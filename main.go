package main

import (
    "fmt"
    "runtime"
    "study/testdeploy/pkg/config"
    "study/testdeploy/pkg/define"
    "study/testdeploy/pkg/models/psql"
    "study/testdeploy/pkg/models/redis"
    "study/testdeploy/pkg/router"
)

// init 初始化配置
func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

// 入口函数
func main() {
    engine, tomlConfig, err := router.InitEngine()
    if err != nil {
        //logging.Errorf("main: 初始化engine出错, err:%v\n", err)
        _ = fmt.Errorf("main: 初始化engine出错, err:%v\n", err)
        return
    }

    // 初始化 DB
    initDB(tomlConfig)

    // 启动服务
    //logging.Infof("run gameconfigapisrv at %v\n", tomlConfig.GetListenAddr())
    fmt.Printf("run gameconfigapisrv at %v\n", tomlConfig.GetListenAddr())
    if err = engine.Run(tomlConfig.GetListenAddr()); err != nil {
        //logging.Errorf("engine run err:%v", err)
        _ = fmt.Errorf("engine run err:%v", err)
    }
}

func initDB(conf *config.Config) {
    // 设置db连接
    psql.NewQipaiDBConn(define.QiPaiDB, conf)

    // 设置Redis连接
    redis.NewVipRedisConn(define.ConfigCache, conf)
}
