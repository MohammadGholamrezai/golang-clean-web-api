package logging

import (
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSingleLogger *zap.SugaredLogger

var logLeverMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
	"panic": zap.PanicLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) GetLogLevel() zapcore.Level {
	level, exists := logLeverMap[l.cfg.Logger.Level]
	if !exists {
		level = zap.DebugLevel
	}

	return level
}

func (l *zapLogger) Init() {
	once.Do(func() {

		w := zapcore.AddSync(zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.cfg.Logger.FilePath,
			MaxSize:    1,
			MaxBackups: 10,
			MaxAge:     5,
			Compress:   true,
			LocalTime:  true,
		}))

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			l.GetLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)).Sugar()
		zapSingleLogger = logger.With("AppName", "MyApp", "LoggerName", "zapLogger")
	})

	l.logger = zapSingleLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *zapLogger) Panic(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := PrepareLogKeys(extra, cat, sub)
	l.logger.Panicw(msg, params...)
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func PrepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}

	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)

	return params
}
