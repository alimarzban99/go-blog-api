package logging

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var once sync.Once

var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	logger *zap.SugaredLogger
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func newZapLogger() *zapLogger {
	logger := &zapLogger{}
	logger.Init()
	return logger
}

func (l *zapLogger) Init() {
	once.Do(func() {
		fileName := fmt.Sprintf("%s%s.%s", config.Config.Logger.Filepath, time.Now().Format("2006-06-01"), "log")
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    1,
			MaxAge:     7,
			LocalTime:  true,
			MaxBackups: 4,
			Compress:   true,
		})

		zapConfig := zap.NewProductionEncoderConfig()
		zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zapConfig),
			w,
			l.getLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel),
		).Sugar()

		zapSinLogger = logger.With("AppName", "MyApp", "LoggerName", "Zaplog")
	})

	l.logger = zapSinLogger
}

func (l *zapLogger) Info(sub SubCategory, msg string) {
	params := l.prepareLogInfo(subCategoryToCategory[sub], sub)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Error(sub SubCategory, msg string) {
	params := l.prepareLogInfo(subCategoryToCategory[sub], sub)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Fatal(sub SubCategory, msg string) {
	params := l.prepareLogInfo(subCategoryToCategory[sub], sub)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exists := zapLogLevelMapping[config.Config.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) prepareLogInfo(cat Category, sub SubCategory) []interface{} {
	extra := make(map[string]interface{})
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return l.logParamsToZapParams(extra)
}

func (l *zapLogger) logParamsToZapParams(keys map[string]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}
