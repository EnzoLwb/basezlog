package cuslog

import (
	"go.uber.org/zap/zapcore"
)

/*
	os.Stdin：标准输入的文件实例，类型为*File
	os.Stdout：标准输出的文件实例，类型为*File
	os.Stderr：标准错误输出的文件实例，类型为*File
*/
type Options struct {
	OutputPaths       []string
	ErrorOutputPaths  []string //错误输出路径 console 或者文件地址
	Level             string   //需要记录的级别
	Formatter         string   //encode json||text
	DisableCaller     bool
	DisableStacktrace bool
	Development       bool //是否为开发模式 开发模式不打印debug级别
	Name              string
	maxAge            int //日志保留天数
	EnableColor       bool
	InitialFields     map[string]interface{} //全局log的额外参数在根对象下 例如 {cluster:"cluster_1"}
	Hook              func(zapcore.Entry, zapcore.SamplingDecision)
}

type OptionFunc func(*Options)

//初始化 全部默认配置
func defaultOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Formatter:         "console",
		EnableColor:       false,
		Development:       false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}

// NewEncoderConfig 设置默认输出文字的规格 暂不支持直接传入Config
func NewEncoderConfig(opts *Options) zapcore.EncoderConfig {
	encodeLevel := zapcore.CapitalLevelEncoder

	// 只有输出控制台时 才有颜色区分
	if opts.Formatter == "console" && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
