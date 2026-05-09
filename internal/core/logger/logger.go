package core_logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	File *os.File
}

type LoggerCongig struct {
	Level string `envconfig:"LEVEL" required:"true"`
	Path  string `envconfig:"PATH" required:"true"`
}

func LoggerFromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value("logger").(*Logger)
	if !ok {
		panic("no logger in context")
	}
	return logger
}

func NewLoggerConfig() (LoggerCongig, error) {
	var c LoggerCongig
	err := envconfig.Process("LOGGER", &c)
	if err != nil {
		err = fmt.Errorf("не получилось пропарсить переменные окружения для конфигуратора логгера:  %w", err)
		return LoggerCongig{}, err
	}
	return c, nil
}

func NewLoggerConfigMust() LoggerCongig {
	config, err := NewLoggerConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func NewLogger(LogConfig LoggerCongig) (*Logger, error) {
	log := zap.NewAtomicLevel()
	if err := log.UnmarshalText([]byte(LogConfig.Level)); err != nil {
		return nil, fmt.Errorf("не получилось создать логгер: %w", err)
	}
	if err := os.MkdirAll(LogConfig.Path, 0755); err != nil {
		return nil, fmt.Errorf("не получилось создать папку для логов: %w", err)
	}

	ts := time.Now().UTC().Format("2006-01-02T15-04-05.00000")
	fp := path.Join(LogConfig.Path, fmt.Sprintf("%s.log", ts))

	logFile, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл на запись логов: %w", err)
	}

	zapConf := zap.NewDevelopmentEncoderConfig()
	zapConf.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.00000")
	zapEncoder := zapcore.NewConsoleEncoder(zapConf)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), log),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), log),
	)
	zapLogger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: zapLogger,
		File:   logFile,
	}, nil
}

func (l *Logger) With(field ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(field...),
		File:   l.File,
	}
}

func (l *Logger) Close() {
	if err := l.File.Close(); err != nil {
		fmt.Println("не удалось закрыть файл логов", err)
	}
}
