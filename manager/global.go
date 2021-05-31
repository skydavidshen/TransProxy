package manager

import (
	"TransProxy/config"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	TP_ROOT_DIR      string
	TP_DB            *gorm.DB
	TP_CACHE_REDIS   *redis.Client
	TP_LOG           *zap.Logger
	TP_CONFIG        *viper.Viper
	TP_SERVER_CONFIG *config.ServerConf
	TP_BASIC_CONFIG  *config.BasicConfig
	TP_MQ_RABBIT     *amqp.Connection
	TP_VALIDATE      *validator.Validate
	TP_ENV           string
)
