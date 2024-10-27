package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	TimeFormat      = "2006-01-02 15:04:05"
	OtPutPaths      = []string{"stderr"} // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
	EncodingConsole = "console"
	EncodingJSON    = "json"
	GroupFlag       = "group"
)

// Field 可以指定关键信息
type Field struct {
	Value interface{}
	Key   string
}

func MakeField(k string, v interface{}) Field {
	return Field{Key: k, Value: v}
}

// Level 声明日志级别
type Level zapcore.Level

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// Logger 定义logger接口
type Logger interface {
	Log() *zap.Logger
	WithGroup(string) Logger
	Debugf(string, ...interface{})
	Info(string, ...Field)
	Infof(string, ...interface{})
	Warn(string, ...Field)
	Warnf(string, ...interface{})
	Error(string, ...Field)
	Errorf(string, ...interface{})
	Fatal(string, ...Field)
	Fatalf(string, ...interface{})
	Println(...interface{})
}

// AppLogger 实现一个logger
type AppLogger struct {
	slg        *zap.SugaredLogger
	lg         *zap.Logger
	level      zap.AtomicLevel
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	group      []string
	MaxSize    int  `json:"max_size"`
	MaxAge     int  `json:"max_age"`
	MaxBackups int  `json:"max_backups"`
	LocalTime  bool `json:"local_time"`
	Compress   bool `json:"compress"`
}

// SetLevel: 维护日志级别的方法
// 使用一个配置接口维护这个方法
func (applg *AppLogger) SetLevel(l Level) {
	applg.level.SetLevel(zapcore.Level(l))
}

func (applg *AppLogger) WithGroup(group string) Logger {
	log := &AppLogger{}
	log.slg = applg.slg
	log.lg = applg.lg
	log.level = applg.level
	if applg.group == nil {
		log.group = make([]string, 0)
	}
	log.group = append(log.group, group)
	return log
}

// GetLevel 获取logger的等级
func (applg *AppLogger) GetLevel() string {
	return applg.level.Level().CapitalString()
}
func (applg *AppLogger) Log() *zap.Logger {
	return applg.lg
}

// Debug logs a message at DebugLevel with SugaredLogger
func (applg *AppLogger) Debugf(temp string, args ...interface{}) {
	applg.slg.Debugf(temp, args...)
}

// Info logs a message at InfoLevel with Logger
func (applg *AppLogger) Info(msg string, fields ...Field) {
	fieldsArr := make([]zap.Field, 0, len(fields))
	fieldsArr = append(fieldsArr, zap.Strings(GroupFlag, applg.group))
	for _, field := range fields {
		fieldsArr = append(fieldsArr, zap.Any(field.Key, field.Value))
	}
	applg.lg.Info(msg, fieldsArr...)
}

// Infof logs a message at InfoLevel with SugaredLogger
func (applg *AppLogger) Println(args ...interface{}) {
	applg.slg.Info(args...)
}

// Infof logs a message at InfoLevel with SugaredLogger
func (applg *AppLogger) Infof(temp string, args ...interface{}) {
	applg.slg.Infof(temp, args...)
}

// Warn logs a message at WarnLevel as ErrorLogger with Logger
func (applg *AppLogger) Warn(msg string, fields ...Field) {
	fieldsArr := make([]zap.Field, 0, len(fields))
	fieldsArr = append(fieldsArr, zap.Strings(GroupFlag, applg.group))
	for _, field := range fields {
		fieldsArr = append(fieldsArr, zap.Any(field.Key, field.Value))
	}
	applg.lg.Warn(msg, fieldsArr...)
}

// Warnf logs a message at WarnLevel with SugaredLogger
func (applg *AppLogger) Warnf(msg string, args ...interface{}) {
	applg.slg.Warnf(msg, args...)
}

// Error logs a message at ErrorLevel with Logger
func (applg *AppLogger) Error(msg string, fields ...Field) {
	fieldsArr := make([]zap.Field, 0, len(fields))
	fieldsArr = append(fieldsArr, zap.Strings(GroupFlag, applg.group))
	for _, field := range fields {
		fieldsArr = append(fieldsArr, zap.Any(field.Key, field.Value))
	}
	applg.lg.Error(msg, fieldsArr...)
}

// Errorf logs a message at ErrorLevel with SugaredLogger
func (applg *AppLogger) Errorf(msg string, args ...interface{}) {
	applg.slg.Errorf(msg, args...)
}

// Fatal logs a message and exit at FatalLevel with Logger
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled
func (applg *AppLogger) Fatal(msg string, fields ...Field) {
	fieldsArr := make([]zap.Field, 0, len(fields))
	fieldsArr = append(fieldsArr, zap.Strings(GroupFlag, applg.group))
	for _, field := range fields {
		fieldsArr = append(fieldsArr, zap.Any(field.Key, field.Value))
	}
	applg.lg.Fatal(msg, fieldsArr...)
}

// Fatalf logs a message at FatalLevel with SugaredLogger
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled
func (applg *AppLogger) Fatalf(msg string, args ...interface{}) {
	applg.slg.Fatalf(msg, args...)
}

type Option func(conf *zap.Config)

// SetLevel 创建Logger实例时, 设置日志等级
func SetLevel(lv Level) Option {
	return func(conf *zap.Config) {
		conf.Level = zap.NewAtomicLevelAt(zapcore.Level(lv))
	}
}

// NopLogger 一个空的log实例
func NopLogger() *AppLogger {
	lg := zap.NewNop()
	return &AppLogger{slg: lg.Sugar(), lg: lg, level: zap.NewAtomicLevelAt(zapcore.Level(DebugLevel))}
}

// DefaultLogger 返回一个基础的默认logger实例
func DefaultLogger() *AppLogger {
	lg := zap.NewNop()
	return &AppLogger{slg: lg.Sugar(), lg: lg, level: zap.NewAtomicLevelAt(zapcore.Level(DebugLevel))}
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,             // 这里可以指定颜色等级显示标识
		EncodeTime:     zapcore.TimeEncoderOfLayout(TimeFormat), // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,          // 调用时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,              // 调用文件路径
	}
}

// NewLogger 创建logger实例对象:
func NewLogger(mode string, conf LogConfig) (*AppLogger, error) {
	logger := &AppLogger{level: zap.NewAtomicLevelAt(zapcore.InfoLevel)}
	if mode == "prod" {
		logger.Filename = conf.Filename
		logger.MaxSize = conf.MaxSize
		logger.MaxBackups = conf.MaxBackups
		logger.MaxAge = conf.MaxAge
		logger.LocalTime = conf.LocalTime
		logger.Compress = conf.Compress
		logger.production()
	} else {
		logger.def()
	}
	return logger, nil
}

// production 生成环境
func (a *AppLogger) production() {
	core := zapcore.NewTee(
		zapcore.NewCore(NewColorConsole(encoderConfig()), zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zapcore.InfoLevel)),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig()), zapcore.AddSync(
			&lumberjack.Logger{ //日志分片
				Filename:   a.Filename,
				MaxSize:    a.MaxSize,    // MB
				MaxBackups: a.MaxBackups, // 最多保留5个备份
				MaxAge:     a.MaxAge,     // days
				LocalTime:  a.LocalTime,  // 是否使用本地时间
				Compress:   a.Compress,   // 是否压缩
			}), zap.NewAtomicLevelAt(zapcore.InfoLevel)),
	)
	lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	a.slg = lg.Sugar()
	a.lg = lg
}

// def 其他环境
func (a *AppLogger) def() {
	core := zapcore.NewCore(
		NewColorConsole(encoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	lg := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1)) //自动上一层
	a.slg = lg.Sugar()
	a.lg = lg
}
