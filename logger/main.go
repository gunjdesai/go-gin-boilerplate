package logger

import (
	"fmt"
	"os"

	"github.com/gunjdesai/go-gin-boilerplate/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	log *zap.Logger
}

var Log *logger

func init() {

	var err error
	Log, err = load(conf.Config.App.Log.Level)

	if err != nil {
		fmt.Println("Error during Logger Initialisation", err)
		os.Exit(10)
	}

}

func load(level string) (*logger, error) {

	log := logger{}
	lvl := getLogLevel(level)

	if err := log.build(lvl); err != nil {
		return &log, err
	}

	return &log, nil

}

func (l *logger) build(lvl zapcore.Level) error {

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(lvl),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.EpochMillisTimeEncoder,
		},
	}

	logger, err := cfg.Build()

	if err != nil {
		return err
	}

	defer logger.Sync()
	l.log = logger

	return nil

}

func (l *logger) Debug(message string, fields ...zap.Field) {

	go l.log.Debug(message, fields...)

}

func (l *logger) Info(message string, fields ...zap.Field) {

	go l.log.Info(message, fields...)

}

func (l *logger) Warn(message string, fields ...zap.Field) {

	go l.log.Warn(message, fields...)

}

func (l *logger) Fatal(message string, fields ...zap.Field) {

	go l.log.Fatal(message, fields...)

}

func (l *logger) Panic(message string, fields ...zap.Field) {

	go l.log.Panic(message, fields...)

}

func getLogLevel(level string) zapcore.Level {

	var lvl zapcore.Level

	switch level {

	case WARN, WARNING:
		lvl = zap.WarnLevel
	case FATAL:
		lvl = zap.FatalLevel
	case INFO:
		lvl = zap.InfoLevel
	case DEBUG:
		lvl = zap.DebugLevel
	case PANIC:
		lvl = zap.PanicLevel
	default:
		lvl = zap.InfoLevel

	}

	return lvl

}
