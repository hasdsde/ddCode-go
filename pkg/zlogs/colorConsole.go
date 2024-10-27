package logger

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	// ColorConsole 带颜色的console
	ColorConsole = "color_console"
	green        = "\033[32m"
	yellow       = "\033[33m"
	red          = "\033[31m"
	blue         = "\033[94m"
	reset        = "\033[0m"
	// cyan         = "\033[36m"

)

var colorMap = map[zapcore.Level]string{
	zapcore.InfoLevel:    green,
	zapcore.WarnLevel:    yellow,
	zapcore.ErrorLevel:   red,
	zapcore.FatalLevel:   red,
	zapcore.InvalidLevel: red,
	zapcore.DPanicLevel:  red,
	zapcore.DebugLevel:   blue,
}

type colorConsoleEncoder struct {
	*zapcore.EncoderConfig
	zapcore.Encoder
}

// NewColorConsole 新建ColorConsole
func NewColorConsole(cfg zapcore.EncoderConfig) (enc zapcore.Encoder) {
	if cfg.ConsoleSeparator == "" {
		// Use a default delimiter of '\t' for backwards compatibility
		cfg.ConsoleSeparator = "\t"
	}
	return colorConsoleEncoder{
		EncoderConfig: &cfg,
		// 使用默认的的 ConsoleEncoder，可以避免重写一遍 ObjectEncoder 等接口
		// PS：ConsoleEncoder 其实也是利用了自带的 JSONEncoder
		Encoder: zapcore.NewConsoleEncoder(cfg),
	}
}

// EncodeEntry 重写 ConsoleEncoder 的 EncodeEntry

func (c colorConsoleEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (buff *buffer.Buffer, err error) {
	buff2, err := c.Encoder.EncodeEntry(ent, fields) // 利用zap已有的实现
	if err != nil {
		return nil, err
	}
	buff = buffer.NewPool().Get()
	buff.AppendString(colorMap[ent.Level]) // 设置颜色
	buff.AppendString(buff2.String())
	buff.AppendString(reset) // 重置
	return buff, err
}
