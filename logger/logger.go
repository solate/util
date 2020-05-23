package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化Logger
func NewLogger(level, path string) (l *zap.Logger, err error) {
	// 设置日志级别，默认为 info
	zapLevel := zap.NewAtomicLevel()
	if err = zapLevel.UnmarshalText([]byte(level)); err != nil {
		return
	}

	// 初始化config
	var config zap.Config
	if zapLevel.Level() == zap.DebugLevel {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	// 将日志输出到指定路径文件中
	if path != "" {
		config.OutputPaths = []string{path}
		config.ErrorOutputPaths = []string{path}
	}

	config.Encoding = "console"
	config.Level = zapLevel
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	//	encoder.AppendString(t.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02T15:04:05.000Z0700"))
	//}
	return config.Build()
}
