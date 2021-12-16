# Release Note
## v0.8.0
* 支持server，client，producer，consumer四种方式的opentelemetry
* 支持通过环境变量EGO_LOG_WRITER，指定输出file或者stderr
有不兼容版本升级，配套组件需要升级到
* ekafka => v0.4.1
* egorm => v0.3.0
* eredis => v0.3.0
  

## v0.7.1
* 升级gin到v1.7.7。修复问题：https://github.com/gin-gonic/gin/releases/tag/v1.7.7
* 支持TrustedPlatform，默认为空，你可以设置为自己的CDN header头获取客户端ip

## v0.7.0
### Feature
* 增加单元测试，测试覆盖率到60%
* 升级gin到master，通过gin的TrustedPlatform参数，解决gin的x-forward-for client ip问题。
* 升级grpc到v1.42.0，grpc新版本解决了客户端、服务端的一些panic问题，详情请看grpc的release信息。
* 支持grpc新版本的attributes，以及克隆attributes方式
* 支持grpc的SetLoggerV2方法，系统日志里会记录grpc底层日志的一些信息。
* 支持通过治理端口，调用job，适用于k8s通过接口调用服务任务
* 支持通过治理端口，调用fgprof方法，拿到函数的io耗时，详情请看https://github.com/felixge/fgprof
* 支持通过proto生成，error、grpc unittest、http api的插件
### FixBug
* 修复在单元测试下，egoErr为ni仍然可以使用其函数做断言操作
* 修复默认日志通过配置重载的时候，在程序关闭时刻flush日志不生效的问题
* 修复grpc在使用option的时候，由于闭包问题导致拦截器里配置数据不正确

## v0.6.3
* 增加大量单元测试
* websocket支持tls
* 去掉无用的gRPC debug配置
* job支持HTTP调用和链路追踪

## v0.6.2
根据环境变量EGO_LOG_EXTRA_KEYS，自定义日志，通常用于打印自定义Headers/Metadata。如用户ID(X-Ego-Uid)、订单ID(X-Ego-Order-Id)等。
* 配置格式 {key1},{key2},{key3}...，多个 key 之间通过 "," 分割。
* 比如 export EGO_LOG_EXTRA_KEYS=X-Ego-Uid,X-Ego-Order-Id
* 这些扩展的追踪字段会根据配置的 key1、key2、key3 等键名，从 Headers(HTTP) 或 Metadata(gRPC)或Context中查找对应值并打印到请求日志中
  支持自定义字段通过context链路传递，将HTTP、gRPC、gorm、redis等常用中间件中传递，自定义字段例如uid、order_id，并记录到日志中。
  拆分自定义header和context
  
## v0.5.8
* HTTP Client: 支持 Interceptor，增加 metric intercetpor
* HTTP/gRPC Server: Access 日志支持打印 Metadata/Header 和 Payload/Body，可通过 enableAccessInterceptorReq、enableAccessInterceptorRes 配置开启
* HTTP/gRPC Server: 支持通过 EGO_LOG_EXTRA_KEYS 环境变量定义额外的日志字段，这些字段从 Headers/Metadata 中自动采集并追加到 Access 日志中
* HTTP Server: 修复 HTTP Access 日志 uri 字段不显示 query 参数问题
* HTTP Server: 优化部分测试用例

## v0.5.7
* http client support trace
* trace is supported by default
* add health check for grpc and http server
* support custom log by context value

## v0.5.5
* 优化websocket
* 支持http到gRPC的代理
* 升级gin
* 支持grpc在连接出问题时候报错
* 支持业务日志的traceid


## v0.5.4
* 正式环境下，panic时候，文件和命令行双输出。并且将field信息，高亮输出到终端。
* 调试模式下，http，gRPC客户端发送请求，会输出行号，该行号可以直接用goland打开。
* resolver阶段，无锁化判断，attributes信息是否变更
* 修改readme文档，描述客户端和服务端gRPC链路玩法
* 优化gRPC客户端设置metadata信息的性能
* 根据客户端头文件的控制，返回gRPC服务端的CPU利用率
* 支持gRPC流式客户端设置metadata信息
* panic增加配置名