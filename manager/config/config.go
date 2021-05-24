package config

import (
	"TransProxy/config"
	"TransProxy/manager"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"log"
)

func GetConfigCenter() (string, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         manager.TP_BASIC_CONFIG.Nacos.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		NotLoadCacheAtStart: true,
		LogDir:              manager.TP_BASIC_CONFIG.Nacos.LogDir,
		CacheDir:            manager.TP_BASIC_CONFIG.Nacos.CacheDir,
		RotateTime:          "1h",
		MaxAge:              3,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      manager.TP_BASIC_CONFIG.Nacos.IpAddr,
			ContextPath: manager.TP_BASIC_CONFIG.Nacos.ContextPath,
			Port:        manager.TP_BASIC_CONFIG.Nacos.Port,
			Scheme:      manager.TP_BASIC_CONFIG.Nacos.Scheme,
		},
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		})
	if err != nil {
		manager.TP_LOG.Error("create config center client fail",
			zap.String("err", err.Error()))
		return "", err
	}

	dataId := manager.TP_BASIC_CONFIG.Nacos.DataId
	group := manager.TP_BASIC_CONFIG.Nacos.Group
	content, errGetConf := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
	if errGetConf != nil {
		manager.TP_LOG.Error("get config center fail",
			zap.String("err", errGetConf.Error()),
			zap.String("dataId", dataId),
			zap.String("group", group))
		return "", errGetConf
	}

	errListen := client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			loadServerConf(data)
			manager.TP_LOG.Info("config changed",
				zap.String("dataId", dataId),
				zap.String("group", group))
		}})
	if errListen != nil {
		manager.TP_LOG.Error("listen config fail",
			zap.String("err", errListen.Error()))
		return "", errListen
	}
	return content, nil
}

func loadServerConf(configCenter string) {
	var serverConf config.ServerConf
	_ = yaml.Unmarshal([]byte(configCenter), &serverConf)
	manager.TP_SERVER_CONFIG = &serverConf
}

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")            // 设置文件名称（无后缀）
	v.SetConfigType("yaml")              // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath(manager.TP_ROOT_DIR) // 设置文件所在路径
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Fatal error config file:", err)
		return nil
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
	})
	//服务配置文件config.yaml加载入对象
	if errBasicConf := v.Unmarshal(&manager.TP_BASIC_CONFIG); errBasicConf != nil {
		log.Println("Read basic config yaml file failed, errBasicConf:", errBasicConf)
		return nil
	}

	configCenter, errCenter := GetConfigCenter()
	if errCenter != nil {
		log.Println("Read center config yaml file failed, errCenter:", errCenter)
		return nil
	}

	loadServerConf(configCenter)
	return v
}
