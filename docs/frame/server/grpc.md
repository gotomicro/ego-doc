# gRPC服务
## 1 Example
[项目地址](https://github.com/gotomicro/ego/tree/master/examples/server/http)

ego版本：``ego@v1.0.0``

## 2 使用方式
```bash
go get github.com/ego-component/ego
```

## 3 gRPC配置
```go
type Config struct {
	Host                       string        // IP地址，默认0.0.0.0
	Port                       int           // Port端口，默认9002
	Deployment                 string        // 部署区域
	Network                    string        // 网络类型，默认tcp4
	EnableMetricInterceptor    bool          // 是否开启监控，默认开启
	EnableTraceInterceptor     bool          // 是否开启链路追踪，默认开启
	SlowLogThreshold           time.Duration // 服务慢日志，默认500ms
	EnableAccessInterceptorReq bool          // 是否开启记录请求参数，默认不开启
	EnableAccessInterceptorRes bool          // 是否开启记录响应参数，默认不开启
	EnableLocalMainIP          bool          // 自动获取ip地址
}
```

## 4 普通服务
### 4.1 用户配置
```toml
[server.grpc]
  host = "127.0.0.1"
  port = 9002
```

### 4.2 用户代码
配置创建一个 ``grpc`` 的配置项，其中内容按照上文配置进行填写。以上这个示例里这个配置key是``server.grpc``

代码中创建一个 ``gRPC`` 服务， egrpc.Load("{{你的配置key}}").Build()，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``gRPC`` 服务后， 将他添加到 ``ego new`` 出来应用的 ``Serve`` 方法中，之后使用的方法和 ``gRPC`` 就完全一致。

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
		helloworld.RegisterGreeterServer(server.Server, &Greeter{server: server})
		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

type Greeter struct {
	server *egrpc.Component
	helloworld.UnimplementedGreeterServer
}

func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "Hello EGO, I'm " + g.server.Address(),
	}, nil
}
```

## 5 开启链路的服务
### 5.1 用户配置
```toml
[trace.jaeger] # 启用链路的核心配置
  ServiceName = "server"
[server.grpc]
  host = "127.0.0.1"
  port = 9002
```
### 5.2 测试代码
[gRPC直连查看链路id](https://github.com/gotomicro/ego/tree/master/examples/grpc/direct)
#### 5.2.1 服务端链路信息
![image](../../images/trace-server-grpc.png)

#### 5.2.2 客户端链路信息
![image](../../images/trace-client-grpc.png)

## 6 开启服务端详细日志信息
### 6.1 测试代码
[gRPC查看详细信息](https://github.com/gotomicro/ego/tree/master/examples/grpc/direct)

### 6.2 用户配置
```toml
[server.grpc]
  host = "127.0.0.1"
  port = 9002
  enableAccessInterceptorReq=true          # 是否开启记录请求参数，默认不开启
  enableAccessInterceptorRes=true          # 是否开启记录响应参数，默认不开启
```
#### 6.3 服务端详细信息
![image](../../images/server-resp-info.png)

## 7 gRPC获取Header头信息
* app
* x-trace-id
* client-ip
* cpu

### 7.1 服务端获取对端应用名的header信息
前提
* gRPC客户端使用了EGO设置app应用名中间件
```go
func getPeerName(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	val, ok2 := md["app"]
	if !ok2 {
		return ""
	}
	return strings.Join(val, ";")
}
```

### 7.2 服务端获取trace id的header信息
前提：
* gRPC客户端使用了EGO设置trace id中间
* gRPC客户端开启链路
* gRPC服务端开启链路
```toml
[trace.jaeger] # 开启链路
```
```go
// 如果开启了全局链路，可以获取链路id
if opentracing.IsGlobalTracerRegistered() {
    etrace.ExtractTraceID(ctx)
}
```

### 7.3 服务端获取client ip的header信息
```go
func getPeerIP(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	// 从metadata里取对端ip
	if val, ok := md["client-ip"]; ok {
		return strings.Join(val, ";")
	}
	// 从grpc里取对端ip
	pr, ok2 := peer.FromContext(ctx)
	if !ok2 {
		return ""
	}
	if pr.Addr == net.Addr(nil) {
		return ""
	}
	addSlice := strings.Split(pr.Addr.String(), ":")
	if len(addSlice) > 1 {
		return addSlice[0]
	}
	return ""
}
```
### 7.4 服务端获取cpu的header信息
未来用于p2c的负载均衡，未实现

<Vssue title="Server-grpc" />