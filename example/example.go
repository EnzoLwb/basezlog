package main

import (
	"errors"
	"fmt"
	log "github.com/EnzoLwb/cuslog"
	"go.uber.org/atomic"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger配置 也可以不配置 像官方log直接使用并兼容；
	// 配置项不止这些 每个都有默认值，详情见defaultOptions()
	count := &atomic.Int64{}
	opts := &log.Options{
		Level:            "debug",
		Formatter:        "json",
		EnableColor:      false,
		DisableCaller:    true,
		OutputPaths:      []string{"test.log", "test2.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"}, //内部error 非error级别日志
		InitialFields: map[string]interface{}{
			"cluster":  "cluster_1",
			"database": "redis",
		},
		Hook: func(entry zapcore.Entry, decision zapcore.SamplingDecision) {
			fmt.Println("count:", count.Inc(), "msg:", entry.Message)
		},
	}

	// 初始化全局logger
	log.Init(opts)
	defer log.Flush()

	// Debug、Info(with field)、Warnf、Errorw的使用 和zap一样
	log.Debug("This is a debug message")
	log.Info("This is a info message",
		log.Int32("int_key", 10),
		log.String("gopher", "go+php"),
		log.Err(errors.New("panic something")),
	) //Logger

	//sugaredLogger
	log.Warnf("This is a formatted %s message", "warn") //这种format的写法 在ide中是没有颜色提示的，而且性能不是最好，所以不建议此写法。

	//临时添加非结构化的值
	log.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")

}
