# gRPC
## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/examples/grpc)

## gRPC配置
```go
type Config struct {
	Addr                         string        // 连接地址，直连为127.0.0.1:9001，服务发现为etcd:///appname
	BalancerName                 string        // 负载均衡方式，默认round robin
	OnFail                       string        // 失败后的处理方式，panic | error
	DialTimeout                  time.Duration // 连接超时，默认3s
	ReadTimeout                  time.Duration // 读超时，默认1s
	SlowLogThreshold             time.Duration // 慢日志记录的阈值，默认600ms
	Debug                        bool          // 是否开启调试，默认不开启，开启后并加上export EGO_DEBUG=true，可以看到每次请求，配置名、地址、耗时、请求数据、响应数据
	EnableBlock                  bool          // 是否开启阻塞，默认开启
	EnableTraceInterceptor       bool          // 是否开启链路追踪，默认开启
	EnableAppNameInterceptor     bool          // 是否开启传递应用名，默认开启
	EnableTimeoutInterceptor     bool          // 是否开启超时传递，默认开启
	EnableMetricInterceptor      bool          // 是否开启监控，默认开启
	EnableAccessInterceptor      bool          // 是否开启记录请求数据，默认不开启
	EnableAccessInterceptorReply bool          // 是否开启记录响应参数，默认不开启
	EnableAccessInterceptorReq   bool          // 是否开启记录请求参数，默认不开启
}
```

## 用户配置
```toml
[grpc.test]
debug = true # 开启后并加上export EGO_DEBUG=true，可以看到每次grpc请求，配置名、地址、耗时、请求数据、响应数据
addr = "127.0.0.1:9002"  
```

## 优雅的Debug
通过开启``debug``配置和命令行的``export EGO_DEBUG=true``，我们就可以在测试环境里看到请求里的配置名、地址、耗时、请求数据、响应数据
![image](../../images/client-grpc.png)


## 用户代码
配置创建一个 ``grpc`` 的配置项，其中内容按照上文配置进行填写。以上这个示例里这个配置key是``grpc.test``

代码中创建一个 ``gRPC`` 客户端， ``egrpc.Load("key").Build()``，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``gRPC`` 客户端后， 将他添加到你所需要的Client里即可。

```go
package main

import (
	"context"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/elog"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

var grpcComp helloworld.GreeterClient

func invokerGrpc() error {
	grpcConn := egrpc.Load("grpc.test").Build()
	grpcComp = helloworld.NewGreeterClient(grpcConn.ClientConn)
	return nil
}

func callGrpc() error {
	_, err := grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "i am client",
	})
	if err != nil {
		return err
	}

	_, err = grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "error",
	})
	if err != nil {
		return err
	}
	return nil
}
```

<Vssue title="Client-grpc" />
