package cuslog

/*
	文件切割相关
*/
import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"time"
)

func getLogWriter(filename string) zapcore.WriteSyncer {
	//logPath := "./"+filename+".log.%Y%m%d"  time.Now().Format("2006-01-02 15:04:05"))
	logPath := "./" + filename + ".log"
	logf, _ := rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(48*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)
	return zapcore.AddSync(logf)
}
