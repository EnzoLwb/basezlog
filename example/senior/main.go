package main

import (
	log "github.com/EnzoLwb/cuslog"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/*
以下是 zap.http_handler.go 内容的翻译：核心就是动态修改日志级别
// PUT请求改变日志级别。在程序运行时改变
// 在程序运行时改变日志级别是非常安全的。有两种内容类型被支持。
//
// 内容类型: application/x-www-form-urlencoded
//
// 对于这种内容类型，级别可以通过请求主体或查询参数提供。
// 一个查询参数。日志级别的URL编码是这样的。
//
// level=debug
//
// 如果请求正文和查询参数都被指定，则请求正文优先于查询参数。
// 的情况下，请求正文优先于查询参数。
//
// 这种内容类型是curl PUT请求的默认类型。下面是两个
// curl请求的例子，它们都将日志级别设置为调试。
//
// curl -X PUT localhost:8080/log/level?level=debug
// curl -X PUT localhost:8080/log/level -d level=debug
//
// 对于任何其他内容类型，有效载荷应该是JSON编码的
//
// {"级别": "信息"}

// 一个curl请求的例子可以是这样的。
//
// curl -X PUT localhost:8080/log/level -H "Content-Type: application/json" -d '{"level": "debug"}'

*/
func main() {
	var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel) //设置默认的日志级别为info级别
	opts := &log.Options{
		AtomicLevel:   atomicLevel,
		Formatter:     "json",
		EnableColor:   false,
		DisableCaller: true,
		OutputPaths:   []string{"stdout"},
	}

	log.Init(opts)
	defer log.Flush()

	http.HandleFunc("/log/level", atomicLevel.ServeHTTP)
	go func() {
		if err := http.ListenAndServe(":8117", nil); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < 10; i++ {

		log.Debug("This is a debug message")
		log.Info("This is a info message",
			log.Int32("int_key", 10),
		)
		log.Warn("This is a warn message")

		log.Error("Message printed with Error")
		time.Sleep(5 * time.Second)
	}
}
