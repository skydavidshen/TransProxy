package log

import (
	"TransProxy/manager"
	"TransProxy/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line_num",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	writerSyncer := getWriterSyncer()
	logLevel := getLogLevel()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writerSyncer, logLevel)

	if logLevel == zap.DebugLevel || logLevel == zap.ErrorLevel {
		logger = zap.New(core, zap.AddStacktrace(logLevel), zap.AddCaller())
	} else {
		logger = zap.New(core, zap.AddCaller())
	}
	return logger
}

func getWriterSyncer() zapcore.WriteSyncer {
	writerSyncer, err := utils.GetWriteSyncer()
	if err != nil {
		panic(fmt.Errorf("Get Write Syncer Failed err: %s \n", err))
	}
	return writerSyncer
}

func getEncoder() zapcore.Encoder {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	return encoder
}

//Get the level based on the configuration file
func getLogLevel() zapcore.Level {
	var level zapcore.Level
	levelConf := manager.TP_SERVER_CONFIG.Log.Level

	switch levelConf { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}