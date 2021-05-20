# K8S
## 1 Example
[获取K8S信息](https://github.com/gotomicro/ego-component/tree/master/ek8s/examples/kubernetesinfo)
[根据K8S endpoints调用gRPC](https://github.com/gotomicro/ego-component/tree/master/ek8s/examples/kubegrpc)

## 2 K8S配置
```go
type Config struct {
    Addr                    string     // 地址
    Debug                   bool       // 调试信息
    Token                   string     // token信息
    Namespaces              []string   // 命名空间列表
    DeploymentPrefix        string     // 前缀
    TLSClientConfigInsecure bool       // 是否开启tls
}
```

## 3 默认配置
* host: KUBERNETES_SERVICE_HOST 环境变量
* port: KUBERNETES_SERVICE_PORT 环境变量
* token: /var/run/secrets/kubernetes.io/serviceaccount/token 文件路径

## 4 获取K8S信息

## 5 根据K8S信息，调用gRPC
### 5.1 K8S配置
```toml
[k8s]
addr=""
token=""
namespaces=["default"]

[grpc.test]
debug = true # 开启后并加上export EGO_DEBUG=true，可以看到每次grpc请求，配置名、地址、耗时、请求数据、响应数据
addr = "k8s:///test:9090"
#balancerName = "round_robin" # 默认值
#dialTimeout = "1s" # 默认值
#enableAccessInterceptor = true
#enableAccessInterceptorRes = true
#enableAccessInterceptorReq = true
```
### 5.2 用户代码
配置创建一个 ``k8s`` 的配置项，其中内容按照上文配置进行填写。以上这个示例里这个配置key是``k8s``

代码中创建一个 ``k8s`` 客户端， ``ek8s.Load("k8s").Build()``，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``k8s`` 客户端后， 将他添加到你所需要的Registry里即可。

```go
// 构建k8s registry，并注册为grpc resolver
registry.DefaultContainer().Build(registry.WithClient(ek8s.Load("k8s").Build()))
// 构建gRPC.ClientConn组件
grpcConn := egrpc.Load("grpc.test").Build()
// 构建gRPC Client组件
grpcComp := helloworld.NewGreeterClient(grpcConn.ClientConn)
fmt.Printf("client--------------->"+"%+v\n", grpcComp)
```