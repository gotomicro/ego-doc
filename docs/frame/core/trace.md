# 链路

## 链路传递
要想使用链路的特性，那么我们一定要会在代码中传递好链路的`Context`。一旦没有传递好`Context`，那么链路数据将会丢失。

以下举常用链路传递方式
### HTTP服务
在`HTTP`服务中的链路的`Context`指的是`c.Request.Context()`。大部分的新手第一次在使用`gin`框架，会将`gin.Context`作为链路`Context`传递下去，而这个用法是错误的用法。以下举例一个在HTTP服务里调用gRPC传递`Context`方式。
```go
func UserInfo(c *gin.Context) {
    reply, err := invoker.UserSvc.UserInfo(c.Request.Context(), &UserInfoReq{
        Uid: c.Query("uid"),
    })
    ...
}
```

### gRPC服务
在`gRPC`服务中的链路的`Context`就是`gRPC`函数里的`Context`，他是可以直接传递的
```go
func (*Order) OrderInfo(ctx context.Context, req *OrderInfoReq) (*OrderInfoReply, error) {
    ...
    reply, err := invoker.UserSvc.UserInfo(ctx, &UserInfoReq{
        Uid: c.Query("uid"),
    })
    ...
}
```

### Gorm客户端
`MySQL`客户端中的链路传递`Context`，如下所示
```go
func (*Order) OrderInfo(ctx context.Context, req *OrderInfoReq) (*OrderInfoReply, error) {
    ...
    var user User
    invoker.Db.WithContext(ctx).Find(&user)
    ...
}
```

### Redis客户端
`MySQL`客户端中的链路传递`Context`，如下所示
```go
func (*Order) OrderInfo(ctx context.Context, req *OrderInfoReq) (*OrderInfoReply, error) {
    ...
    invoker.Redis.Get(ctx, "hello")
    ...
}
```

### HTTP客户端
resty没有很好的拦截器，不太好封装链路，所以只能自行调用
```go
tracer := etrace.NewTracer(trace.SpanKindClient)
req := httpComp.R()
ctx, span := tracer.Start(context.Background(), "callHTTP()", propagation.HeaderCarrier(req.Header))
defer span.End()

fmt.Println(span.SpanContext().TraceID())
info, err := req.SetContext(ctx).Get("http://127.0.0.1:9007/hello")
if err != nil {
    return err
}
fmt.Println(info)
```

## 业务记录链路id日志
在ego里面提供了链路id的函数，如下所示
```go
// FieldCtxTid 设置链路id
func FieldCtxTid(ctx context.Context) Field {
	return String("tid", etrace.ExtractTraceID(ctx))
}
```
业务记录日志只需要如下即可
### HTTP服务
在`HTTP`服务中的链路的`Context`指的是`c.Request.Context()`。大部分的新手第一次在使用`gin`框架，会将`gin.Context`作为链路`Context`传递下去，而这个用法是错误的用法。以下举例一个在HTTP服务里调用gRPC传递`Context`方式。
```go
func UserInfo(c *gin.Context) {
    reply, err := invoker.UserSvc.UserInfo(c.Request.Context(), &UserInfoReq{
        Uid: c.Query("uid"),
    })
    if err != nil {
        elog.Error("请求用户服务错误",elog.FieldCtxTid(c.Request.Context()),elog.FieldErr(err))
        return
    }
    ...
}
```

### grpc服务
```go
func (*Order) OrderInfo(ctx context.Context, req *OrderInfoReq) (*OrderInfoReply, error) {
    ...
    reply, err := invoker.UserSvc.UserInfo(ctx, &UserInfoReq{
        Uid: c.Query("uid"),
    })
    if err != nil {
        elog.Error("请求用户服务错误",elog.FieldCtxTid(ctx),elog.FieldErr(err))
        return
    }
    ...
}
```

## 链路trace id
框架默认自动开启链路，所有服务端或者客户端`access`日志会自动加入`trace id`，用户无需关心。 结合你的前端框架，报错时候，可以弹出链路id。

![](../../images/trace.png)


## 自定义链路
在应用的环境变量里，配置你需要在日志里添加自定义属性的环境变量`export EGO_LOG_EXTRA_KEYS=X-Ego-Uid,X-Order-Id`，框架在读到这个环境变量里`EGO_LOG_EXTRA_KEYS`，会根据逗号分割，解析用户配置了多少个自定义的链路属性，后续的`Ego`组件会根据这个属性，自动将组件日志里追加这些用户自定义属性。

### 如何在HTTP服务中鉴权后，在链路中添加`Uid`信息
```go
import "github.com/gotomicro/ego/core/transport"

func checkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Access-Token")
		reply := checkToken(accessToken)
		// 使用ego中的transport赋值，会将X-Ego-Uid转为ego中context key的结构体
		// 避免ego内置context key和业务方自己的context key冲突
		parentContext := transport.WithValue(c.Request.Context(), "X-Ego-Uid", reply.Uid)
		c.Request = c.Request.WithContext(parentContext)
		c.Next()
	}
}
```

### 通用Context的赋值和取值
```go
import "github.com/gotomicro/ego/core/transport"

func main() {
	// 因为这里是模拟的自定义属性，为了测试，手动写入X-Ego-Uid
	// 实际业务应用启动后，该数据会从环境变量里获取
    transport.Set([]string{"X-Ego-Uid"})
    // 赋值
    ctx := transport.WithValue(context.Background(), "X-Ego-Uid", 9527)
    // 取值
    value := transport.Value(ctx, "X-Ego-Uid")
    fmt.Println(value)
}
```

### 最终达到效果
[链路数据流向原理链接](https://ego-org.com/micro/chapter2/trace.html)

例如`MySQL`的`access`日志中记录了`uid`和`trace id`信息。
![img.png](../../images/egoaccesstrace3.png)

以上就是Ego将日志中实现了链路的数据流向。快来体验吧。

体验版本为：
* github.com/gotomicro/ego@v0.6.2
* github.com/gotomicro/ego-component/egorm@v0.2.2
* github.com/gotomicro/ego-component/eredis@v0.2.3

