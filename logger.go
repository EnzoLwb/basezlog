package cuslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

var std = New() //创建全局默认的logger

type logger struct {
	opt *options
	log   *zap.SugaredLogger
	mu sync.Mutex
}

func New(opts ...OptionFunc) *logger {
	//初始化选项
	opt := initOptions(opts...)

	//日志将写到哪里去
	writeSyncer := zapcore.AddSync(opt.output) // file, _ := os.Create("./test.log")

	//如何写入日志 JSON || TEXT
	encoder := getJsonEncoder(opt)

	// core 的三个配置：编码器 写入文件地址 日志级别
	core := zapcore.NewCore(encoder, writeSyncer, opt.level)

	//将调用方栈信息记录下来
	zapLog := zap.New(core,zap.AddCaller())

	//实例化全局logger
	logger := &logger{
		opt: opt,
		log: zapLog.Sugar(),
	}
	return logger
}

// getJsonEncoder 设置 编码器的信息
func getJsonEncoder(opts *options) zapcore.Encoder {
	config := NewEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder //修改时间编码器
	config.EncodeLevel = zapcore.CapitalLevelEncoder //在日志文件中使用大写字母记录日志级别
	if opts.formatter == EncodeJson {
		return zapcore.NewJSONEncoder(config)
	}else if  opts.formatter == EncodeText {
		return zapcore.NewConsoleEncoder(config)
	}else {
		log.Fatalf("options.formatter is only json or text")
		return nil
	}
}

// NewEncoderConfig 设置默认输出文字的规格
func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}