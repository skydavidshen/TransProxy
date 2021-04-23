package manager

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	TP_ROOT_DIR string
	TP_DB       *gorm.DB
	TP_REDIS    *redis.Client
	TP_LOG      *zap.Logger
	TP_CONFIG   *viper.Viper
)
