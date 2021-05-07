package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	集成 zap logger

 */

type ZapConfig struct {

	Development   bool

}

// NewZapLogger new zapLogger and replace global logger
func NewZapLogger(config *ZapConfig) (*zap.Logger, error) {
	var zapConfig zap.Config

	if config.Development {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	// 替换zap包全局logger
	undo1 := zap.ReplaceGlobals(logger)
	defer undo1()
	// 输出到std
	undo2 := zap.RedirectStdLog(logger)
	defer undo2()

	return logger, nil
}
