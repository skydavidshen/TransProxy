package utils

import (
	"com.pippishen/trans-proxy/manager"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"path"
	"time"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	logDir := manager.TP_CONFIG.Get("log.director").(string)
	rotateLogs, error := zaprotatelogs.New(
		path.Join(logDir, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(30*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.AddSync(rotateLogs), error
}