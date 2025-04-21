package logging

import (
	"os"
	"sync"

	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var zeroSingleLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLeverMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLeverMap[l.cfg.Logger.Level]
	if !exists {
		level = zerolog.DebugLevel
	}

	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Could not open log file: " + err.Error())
		}

		logger := zerolog.New(file).With().Timestamp().Str("AppName", "MyApp").Str("source", "zeroLogger").Logger()
		zerolog.SetGlobalLevel(l.getLogLevel())
		zeroSingleLogger = &logger
	})

	l.logger = zeroSingleLogger
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.
		Debug().
		Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.
		Warn().
		Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.
		Error().
		Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.
		Fatal().
		Msgf(template, args...)
}

func (l *zeroLogger) Panic(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.
		Panic().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Panicf(template string, args ...interface{}) {
	l.logger.
		Panic().
		Msgf(template, args...)
}
