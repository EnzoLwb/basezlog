package cuslog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

//两种创建方式 一种是New defaultOptions 另一种是 Init(&Options实例)
var (
	std = New(defaultOptions())
	mu  sync.Mutex
)

type logger struct {
	zapLogger *zap.Logger
}

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = New(opts)
}

func New(opts *Options) *logger {
	if opts == nil {
		opts = defaultOptions()
	}

	loggerConfig := InitZapConfig(opts)
	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel),
		zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	//实例化全局logger
	logger := &logger{
		zapLogger: l,
	}
	return logger
}

//InitZapConfig 类似于zap的NewProduction NewDevelopment方法
func InitZapConfig(opts *Options) *zap.Config {
	//将字符串level转化为zap level ，使用zap提供的转换方法
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	return &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		// 采样设置，记录全局的CPU、IO负载 hook
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
			Hook:       opts.Hook,
		},
		// 支持编码json或者console 咱以后扩展
		Encoding:         opts.Formatter,
		EncoderConfig:    NewEncoderConfig(opts),
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
		// 可以额外添加的参数
		InitialFields: opts.InitialFields,
	}
}

//格式化 EncoderConfig中的字段
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

//格式化 EncoderConfig中的字段
func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

// Flush zap 的 Sync()
func Flush() { std.Flush() }

func (l *logger) Flush() {
	_ = l.zapLogger.Sync()
}
