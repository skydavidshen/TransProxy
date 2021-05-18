package config

import (
	"TransProxy/manager"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const configPath string = "config.yaml"

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config") // 设置文件名称（无后缀）
	v.SetConfigType("yaml")   // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath(manager.TP_ROOT_DIR)  // 设置文件所在路径
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
	})

	//服务配置文件config.yaml加载入对象
	if err := v.Unmarshal(&manager.TP_SERVER_CONFIG); err != nil {
		manager.TP_LOG.Error("Read server config yaml file failed, err:",
			zap.String("err", err.Error()))
	}

	return v
}