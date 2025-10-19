package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	sl *zap.SugaredLogger
}

func NewLogger() *ZapLogger {
	config := zap.NewProductionConfig()

	// options zap
	config.OutputPaths = []string{"stdout", "logs/app.log"}
	config.Level.SetLevel(zapcore.InfoLevel)
	config.Encoding = "json"

	rawLogger, err := config.Build()
	if err != nil {
		return nil
	}

	return &ZapLogger{sl: rawLogger.Sugar()}
}

func (l *ZapLogger) Infof(formatStr string, args ...any) {
	l.sl.Infof(formatStr, args...)
}

func (l *ZapLogger) Debugf(formatstr string, args ...any) {
	l.sl.Debugf(formatstr, args...)
}

func (l *ZapLogger) Errorf(formatstr string, args ...any) {
	l.sl.Errorf(formatstr, args...)
}

func (l *ZapLogger) Warnf(formatStr string, args ...any) {
	l.sl.Warnf(formatStr, args...)
}

func (l *ZapLogger) Fatalf(formatStr string, args ...any) {
	l.sl.Fatalf(formatStr, args...)
}

func InitLogger() ZapLogger {
	return *NewLogger()
}
