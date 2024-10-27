package logger

type LogConfig struct {
	Level      string `json:"level" yaml:"Level"`
	Filename   string `json:"filename" yaml:"Filename"`
	MaxSize    int    `json:"maxSize" yaml:"MaxSize"`
	MaxAge     int    `json:"maxAge" yaml:"MaxAge"`
	MaxBackups int    `json:"maxBackups" yaml:"MaxBackups"`
	LocalTime  bool   `json:"localTime" yaml:"LocalTime"`
	Compress   bool   `json:"compress" yaml:"Compress"`
}
