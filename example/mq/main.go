package main

import (
	log "github.com/EnzoLwb/cuslog"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var _ io.Writer = (*redisWriter)(nil)
var _ io.Writer = (*RabbitMqQueueWriter)(nil)

var key = "log_list"

func main() {
	opts := &log.Options{
		Level:         "debug",
		Formatter:     "json",
		EnableColor:   false,
		DisableCaller: true,
		OutputPaths:   []string{"stdout"},
	}
	var Cores []zapcore.Core

	// 生成一个console core
	encoder := zapcore.NewJSONEncoder(log.NewEncoderConfig(opts))
	stdCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.InfoLevel)

	// 生成一个redis core
	cli := redis.NewClient(&redis.Options{
		Addr: "192.168.105.114:6379",
		//Password: "123456",
	})
	_, err := cli.Ping().Result()
	if err != nil {
		panic(err)
	}
	writer := NewRedisWriter(key, cli)
	redisCore := zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.InfoLevel)

	//生成一个 rabbitmq core
	conn, err := amqp.Dial("amqp://admin:admin@192.168.105.114:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		key,      // 使用命名的交换器
		"fanout", // 交换器类型
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	mqWriter := NewRabbitMqWriter(key, ch)
	rmqCore := zapcore.NewCore(encoder, zapcore.AddSync(mqWriter), zap.InfoLevel)

	//append Core
	Cores = append(Cores, stdCore, redisCore, rmqCore)
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
