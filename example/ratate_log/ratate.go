package main

/*
	文件切割、不同level输出到不同日志
*/
import (
	log "github.com/EnzoLwb/cuslog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	opts := &log.Options{
		Level:         "debug",
		Formatter:     "json",
		EnableColor:   false,
		DisableCaller: true,
		OutputPaths:   []string{"stdout"},
	}

	var Cores []zapcore.Core
	encoder := zapcore.NewJSONEncoder(log.NewEncoderConfig(opts))

	Cores = append(Cores, zapcore.NewCore(encoder, getLogWriter("infolog"), zap.InfoLevel))
	Cores = append(Cores, zapcore.NewCore(encoder, getLogWriter("warnlog"), zap.WarnLevel))
	log.NewByCores(Cores)
	defer log.Flush()

	//测试
	log.Debug("This is a debug message")
	log.Info("This is a info message",
		log.Int32("int_key", 10),
	)
	log.Warn("This is a warn message")

	log.Error("Message printed with Error")
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	logPath := "./" + filename + ".log"
	logf, _ := rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(48*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)
	return zapcore.AddSync(logf)
}
