package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Conf *AppConfig
)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	*MysqlConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// Init 初始化配置
func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/config.yaml")

	// 查找并读取配置文件
	err = viper.ReadInConfig()
	// 处理读取配置文件的错误
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//解析到结构体Conf
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
