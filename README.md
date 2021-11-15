##### `EnzoLwb/cuslog` 是个开箱即用日志包，基于 zap 包封装。具有如下特性：
 - [ ] 能够将事件记录到文件中，而不是应用程序控制台
 - [ ] 日志切割：能够根据文件大小、时间或间隔等来切割日志文件
 - [ ] 能够打印基本信息，如调用文件/函数名和行号，日志时间等
 - [ ] 支持不同的日志级别。例如INFO，DEBUG，ERROR等（官方log只有print、fatal、panic级别
 - [ ] 兼容官方log
 - [ ] 结构化打印 方便Filebeat、Logstash Shipper等工具记录
 - [ ] zap的hook功能
 - [ ] 配置颜色输出
 - [ ] 日志投递：投递到 Elasticsearch、Kafka 等组件
 - [ ] zap动态开关日志级别：无需重启服务，通过请求改变日志级别

### Usage

#### 开箱即用


#### 配置选项后使用


#### 更多功能介绍请查看example.go

