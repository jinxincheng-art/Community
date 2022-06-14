package main

import (
	"community/config"
	"community/dao"
	"community/logger"
	"community/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	//初始化配置
	if err := config.Init(); err != nil {
		fmt.Printf("init config failed err: %v\n", zap.Error(err))
		return
	}

	//初始化日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger failed err : %v\n", zap.Error(err))
		return
	}

	//初始化数据库
	if err := dao.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed err : %v\n", zap.Error(err))
		return
	}

	//初始化路由
	if err := routes.InitRouter(gin.Default()); err != nil {
		fmt.Printf("init roters failed err : %v\n", zap.Error(err))
		return
	}
}
