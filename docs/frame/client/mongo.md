# Mongo
## 1 Example
[项目地址](https://github.com/gotomicro/ego-component/tree/master/emongo/examples)

ego版本：``ego@v0.5.3``

## 2 mongo配置
```go
type Config struct {
    DSN                        string        // DSN DSN地址
    Debug                      bool          // Debug 是否开启debug模式
    SocketTimeout              time.Duration // SocketTimeout 创建连接的超时时间
    PoolLimit                  int           // PoolLimit 连接池大小(最大连接数)
    EnableMetricInterceptor    bool          // EnableMetricInterceptor 是否启用prometheus metric拦截器
    EnableAccessInterceptorReq bool          // EnableAccessInterceptorReq 是否启用access req拦截器，此配置只有在EnableAccessInterceptor=true时才会生效
    EnableAccessInterceptorRes bool          // EnableAccessInterceptorRes 是否启用access res拦截器，此配置只有在EnableAccessInterceptor=true时才会生效
    EnableAccessInterceptor    bool          // EnableAccessInterceptor 是否启用access拦截器
    SlowLogThreshold           time.Duration // SlowLogThreshold 慢日志门限值，超过该门限值的请求，将被记录到慢日志中
}
```

## 3 优雅的Debug
通过开启``debug``配置和命令行的``export EGO_DEBUG=true``，我们就可以在测试环境里看到请求里的配置名、地址、耗时、请求数据、响应数据
![img.png](../../images/frame/client/mongo/mongo1.png)


## 4 用户配置
```toml
[mongo]
debug=true
dsn="mongodb://user:password@localhost:27017,localhost:27018"
```

## 5 用户代码
```go
var stopCh = make(chan bool)
	// 假设你配置的toml如下所示
	conf := `
[mongo]
	debug=true
	dsn="mongodb://user:password@localhost:27017,localhost:27018"
`
	// 加载配置文件
err := econf.LoadFromReader(strings.NewReader(conf), toml.Unmarshal)
if err != nil {
    panic("LoadFromReader fail," + err.Error())
}

// 初始化emongo组件
cmp := emongo.Load("mongo").Build()
coll := cmp.Client.Database("test").Collection("cells")
findOne(coll)

stopCh <- true
```