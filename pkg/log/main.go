package log

import (
	"aura-test/pkg/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger use zap.Logger for log instance
var Logger *zap.Logger

func init() { //nolint
	atom := zap.NewAtomicLevelAt(zapcore.Level(config.EnvForge().GetInt("LOG_LEVEL")))
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		atom,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.Development())
	defer Logger.Sync() // nolint

	zap.ReplaceGlobals(Logger)
}

// Error uses zap.sugar to construct and log a message.
func Error(err ...any) {
	Logger.Sugar().Error(err...)
}

// Debug uses zap.sugar to construct and log a message.
func Debug(err ...any) {
	Logger.Sugar().Debug(err...)
}

// Info uses zap.sugar to construct and log a message.
func Info(err ...any) {
	Logger.Sugar().Info(err...)
}

// Warning uses zap.sugar to construct and log a message.
func Warning(err ...any) {
	Logger.Sugar().Warn(err...)
}

// Fatal uses zap.sugar to construct and log a message.
func Fatal(err ...any) {
	Logger.Sugar().Fatal(err...)
}

// Panic uses zap.sugar to construct and log a message.
func Panic(err ...any) {
	Logger.Sugar().Panic(err...)
}

// WithFields adds a variadic number of fields to the logging context
func WithFields(args ...any) *zap.SugaredLogger {
	return Logger.Sugar().With(args...)
}
