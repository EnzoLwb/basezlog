package cuslog

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type options struct {
	output        io.Writer //日志文件路径
	level         zapcore.Level //需要记录的级别
	stdOut        bool //是否输出到终端
	formatter     Formatter //encode json||text
	disableCaller bool
	maxAge 		  int //日志保留天数
	enableColor   bool
}

type OptionFunc func(*options)

//默认初始化
func initOptions(opts ...OptionFunc) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}
	fmt.Println("o\n")
	fmt.Println(o)
	//不填写就输出到控制台
	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == "" {
		o.formatter = EncodeJson
	}
	return
}

func WithOutput(output io.Writer) OptionFunc {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level zapcore.Level) OptionFunc {
	return func(o *options) {
		o.level = level
	}
}

func WithFormatter(formatter Formatter) OptionFunc {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(disableCaller bool) OptionFunc {
	return func(o *options) {
		o.disableCaller = disableCaller
	}
}

func WithStdOut(stdOut bool) OptionFunc {
	return func(o *options) {
		o.stdOut = stdOut
	}
}


// cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
func (l *logger) SetOptions(opts ...OptionFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}
