package log

import (
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_runningAtomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	_logger             atomic.Value
	defaultLogger       *zap.Logger
	defaultLoggerInit   sync.Once
)

type Field = zap.Field

// Log represents a logging configuration.
type Config struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

func init() {
	defaultLogger = DefaultLogger()
}

// sDefaultLogger creates a default logger for uninitialized usage.
func DefaultLogger() *zap.Logger {
	logConf := zap.NewProductionConfig()
	logConf.DisableStacktrace = true
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l, err := logConf.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return l
}

// InitLogger initializes a zap logger from user config.
func InitLogger(cfg Config) error {
	if err := initializeLogger(cfg); err != nil {
		return err
	}
	return nil
}

// initLogger initializes a zap logger for different module.
func initializeLogger(cfg Config) error {
	// parse logging level
	if err := _runningAtomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}
	conf := zap.NewProductionConfig()
	conf.Level = _runningAtomicLevel
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.DisableStacktrace = true
	// mode
	if cfg.Mode == "console" {
		conf.Encoding = "console"
		conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	l, err := conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	_logger.Store(l)
	defaultLogger = l
	return nil
}

func getLogger() *zap.Logger {
	defaultLoggerInit.Do(func() {
		InitLogger(Config{Level: "info"})
	})
	return defaultLogger
}

func Info(msg string, fields ...Field) {
	getLogger().Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	getLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	getLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	getLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	getLogger().Fatal(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	getLogger().Sugar().Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	getLogger().Sugar().Errorf(template, args...)
}

// Other methods follow similarly...

func String(key string, val string) Field {
	return zap.String(key, val)
}

func Uint16(key string, val uint16) Field {
	return zap.Uint16(key, val)
}

func Uint32(key string, val uint32) Field {
	return zap.Uint32(key, val)
}

func Uint64(key string, val uint64) Field {
	return zap.Uint64(key, val)
}

func Stack() Field {
	return zap.Stack("stack")
}

func Reflect(key string, val interface{}) Field {
	return zap.Reflect(key, val)
}

func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

func Int32(key string, val int32) Field {
	return zap.Int32(key, val)
}

func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

