package manager

import (
	"com.pippishen/trans-proxy/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	TP_ROOT_DIR       string
	TP_DB             *gorm.DB
	TP_CACHE_REDIS    *redis.Client
	TP_LOG            *zap.Logger
	TP_CONFIG         *viper.Viper
	TP_SERVER_CONFIG  *config.ServerConf
	TP_MQ_RABBIT      *amqp.Connection
)
