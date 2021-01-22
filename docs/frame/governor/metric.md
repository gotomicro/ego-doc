## 监控信息
### 服务端标准
以下是框架记录的指标，prometheus在采集时加入指标，例如job或者app等数据，所以可以由他来过滤一些应用信息数据
#### server端计数器
* server_handle_total

|名称|描述|用法|
| --- | --- | --- |
|type|类型|
|method|方法|
|peer|对端节点|
|code|状态码|

type 有三种类型
* http
* unary
* stream

method 有两种类型
* HTTP的method是 c.Request.Method+"."+c.Request.URL.Path
* gRPC的method是 grpc.UnaryServerInfo.FullMethod

peer 取数据
* HTTP的peer是从header头里取出app，这个为对端节点的应用名
* gRPC的peer是从header头里取出app，这个为对端节点的应用名

code 取数据
* HTTP的code是HTTP状态码
* gRPC的code是gRPC返回的Message，只记录系统错误码，系统错误码成功为OK，非系统错误码全部记录为``biz err``，防止prometheus错误爆炸

#### server直方图
server_handle_seconds

|名称|描述|用法|
| --- | --- | --- |
|type|类型|
|method|方法|
|peer|对端节点|

type 有三种类型
* http
* unary
* stream

method 有两种类型
* HTTP的method是 c.Request.Method+"."+c.Request.URL.Path
* gRPC的method是 grpc.UnaryServerInfo.FullMethod

peer 取数据
* HTTP的peer是从header头里取出app，这个为对端节点的应用名
* gRPC的peer是从header头里取出app，这个为对端节点的应用名
