# 日志
## Example
* [终端显示日志](https://github.com/gotomicro/ego/tree/master/example/logger/console)
* [文本显示日志](https://github.com/gotomicro/ego/tree/master/example/logger/file)
* [日志动态修改级别](https://github.com/gotomicro/ego/tree/master/example/logger/watch)


## 日志配置
框架在处理的日志区分为框架日志和业务日志，了解日志，请阅读日志和错误处理。

日志配置的数据结构如下
```go
// Config ...
type Config struct {
	Debug               bool          // 配置文件控制日志输出到终端还是文件，true到终端，false到文件
	Dir                 string        // 日志输出目录，默认logs
	Name                string        // 日志文件名称，默认框架日志mocro.sys，业务日志default.log
	Level               string        // 日志初始等级，默认info级别
	AddCaller           bool          // 是否添加调用者信息，默认不加调用者信息
	MaxSize             int           // 日志输出文件最大长度，超过改值则截断，默认500M
	MaxAge              int           // 日志存储最大时间，默认最大保存天数为7天
	MaxBackup           int           // 日志存储最大数量，默认最大保存文件个数为10个
	Interval            time.Duration // 日志轮转时间，默认1天
	Async               bool          // 是否异步，默认异步
	FlushBufferSize     int           // 缓冲大小，默认256 * 1024B
	FlushBufferInterval time.Duration // 缓冲时间，默认5秒
	Fields              []zap.Field   // 日志初始化字段
}
```

## 终端显示日志
在运行程序前开启环境变量 ``EGO_DEBUG=true``，可以把所有日志输出到终端。并且开启了该指令后，日志的时间变成 ``time.Time`` 数据结构。

```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
)

//  export EGO_DEBUG=true && go run main.go
func main() {
	err := ego.New().Invoker(func() error {
		elog.Info("logger info", elog.String("gopher", "ego"), elog.String("type", "command"))
		return nil
	}).Run()
	if err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}
```

## 文件显示日志
当``EGO_DEBUG`` 环境变量不存在或者``EGO_DEBUG=false``的时候，日志默认输出到 logs 目录下。
```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
)

//  export EGO_DEBUG=false && go run main.go
func main() {
	err := ego.New().Invoker(func() error {
		elog.Info("logger info", elog.String("gopher", "ego"), elog.String("type", "command"))
		return nil
	}).Run()
	if err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}
```

## 动态日志级别
框架里自带的框架日志和业务日志都默认支持动态更改日志级别。当程序启动后，你可以在配置文件里更改lv的级别从info改为debug，就可以看到动态生效的debug日志，该方法非常利于研发排查线上问题，倡导大家线下多打debug日志，线上用info级别日志，出现线上问题可以改变日志级别，快速排查问题。

```toml
[logger.default]
    level = "info"
```

```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"time"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New(ego.WithHang(true)).Invoker(func() error {
		go func() {
			for {
				elog.Info("logger info", elog.String("gopher", "ego1"), elog.String("type", "file"))
				elog.Debug("logger debug", elog.String("gopher", "ego2"), elog.String("type", "file"))
				time.Sleep(1 * time.Second)
			}
		}()
		return nil
	}).Run()
	if err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}
```

## 日志字段
EGO的字段是确定类型的，通过正交查询方式，减少索引字段个数，同时方便创建索引。后续字段类型尽量像[opentrace](https://github.com/opentracing-contrib/opentracing-specification-zh/blob/master/semantic_conventions.md)靠拢

|名称|类型|描述|
| --- | --- | --- |
|lv|string|日志级别|
|ts	|string|时间戳|
|msg|string|日志信息|
|app|string|应用名称|
|iid|string|应用实例id|
|tid|string|请求trace id|
|color|string|染色|
|comp|string|类库或组件。如 "grpc", "http", "redis".|
|compName|string|组件配置key作为唯一标识|
|addr|	string|	依赖的实例名称。以mysql为例，"dsn = "root:root@tcp(127.0.0.1:3306)/ego?charset=utf8"，addr为 "127.0.0.1:3306".|
|cost|	int|	耗时时间|
|code|	int|	用户侧响应的状态码|
|meth|	string|	对于redis是command、对于http是url、对于mysql是sql|
|host|	string|	主机名|
|ip|	string|	主机IP|
|peerApp|	string|	对端应用名称|
|peerHost|	string|	对端主机名|
|errKind|	string|	错误类型，用于收敛|
|err|	string|	错误信息|

## 日志
* 慢日志
* 错误日志