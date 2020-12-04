# gRPC服务
## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/example/server/http)

## HTTP配置
```
type Config struct {
	Host                    string // IP地址，默认0.0.0.0
	Port                    int    // Port端口，默认9002
	Deployment              string // 部署区域
	Network                 string // 网络类型，默认tcp4
	DisableTrace            bool   // 禁用监控，默认否
	DisableMetric           bool   // 禁用trace，默认否
	SlowLogThresholdInMilli int64  // 服务慢日志，默认500ms
}
```

## 用户配置
```
[server.grpc]
  host = "127.0.0.1"
  port = 9002
```

## 用户代码
配置创建一个 ``{{你的配置key}}`` 的配置项，其中内容按照上文HTTP的配置进行填写。以上这个示例里这个配置key是``server.grpc``

代码中创建一个 ``gRPC`` 服务， egrpc.Load("{{你的配置key}}").Build()，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``http`` 服务后， 将他添加到 ``ego new`` 出来应用的 ``Serve`` 方法中，之后使用的方法和 ``gRPC`` 就完全一致。

```go
package main

import (
	"context"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egrpc"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(func() server.Server {
		server := egrpc.Load("server.grpc").Build()
		helloworld.RegisterGreeterServer(server.Server, &Greeter{})
		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

type Greeter struct {
	server *egrpc.Component
}

func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "Hello EGO, I'm " + g.server.Address(),
	}, nil
}
```