# ETCD服务注册与发现使用
## 1 配置
### 1.1 Registry配置
```go
type Config struct {
	Scheme       string        // 协议
	Prefix       string        // 注册前缀
	ReadTimeout  time.Duration // 读超时
	ServiceTTL   time.Duration // 服务续期
	OnFailHandle string        // 错误后处理手段，panic，error
}
```
### 1.2 ETCD配置
```go
type config struct {
	Addrs                        []string      // 地址
	CertFile                     string        // cert file
	KeyFile                      string        // key file
	CaCert                       string        // ca cert
	UserName                     string        // 用户名
	Password                     string        // 密码
	ConnectTimeout               time.Duration // 连接超时时间
	AutoSyncInterval             time.Duration // 自动同步member list的间隔
	EnableBasicAuth              bool          // 是否开启认证
	EnableSecure                 bool          // 是否开启安全
	EnableBlock                  bool          // 是否开启阻塞，默认开启
	EnableFailOnNonTempDialError bool          // 是否开启gRPC连接的错误信息
}
```

## 2 服务注册
### 2.1 服务注册配置
```toml
[server.grpc]
    port = 9003
[etcd]
    addrs=["127.0.0.1:2379"]
    connectTimeout = "1s"
    secure = false
[registry]
    scheme = "etcd" # grpc resolver默认scheme为"etcd"，你可以自行修改
    #serviceTTL = "10s"
```

### 2.2 服务注册代码
```go
package main

import (
	"context"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/examples/helloworld"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().
		Registry(registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))).
		Serve(func() server.Server {
			server := egrpc.Load("server.grpc").Build()
			helloworld.RegisterGreeterServer(server.Server, &Greeter{server: server})
			return server
		}()).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

type Greeter struct {
	server *egrpc.Component
}

func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if request.Name == "error" {
		return nil, status.Error(codes.Unavailable, "error")
	}

	return &helloworld.HelloReply{
		Message: "Hello EGO, I'm " + g.server.Address(),
	}, nil
}
```


## 3 服务发现
### 3.1 服务发现
```toml
[grpc.test]
debug = true # 开启后并加上export EGO_DEBUG=true，可以看到每次grpc请求，配置名、地址、耗时、请求数据、响应数据
addr = "etcd:///main"
#balancerName = "round_robin" # 默认值
#dialTimeout = "1s" # 默认值
#enableAccessInterceptor = true
#enableAccessInterceptorRes = true
#enableAccessInterceptorReq = true

[etcd]
addrs=["127.0.0.1:2379"]
connectTimeout = "1s"
secure = false

[registry]
scheme = "etcd" # grpc resolver默认scheme为"etcd"，你可以自行修改
```

### 3.2 服务注册代码
```go
package main

import (
	"context"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/elog"

	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/examples/helloworld"
	"github.com/gotomicro/ego-component/eetcd/registry"
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
	// 注册resolver
	registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))
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