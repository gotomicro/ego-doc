# EGO生态里的组件统一了日志，监控指标。
* EGO版本: v1.2.2
* Egorm版本: v1.1.3
* Eredis版本: v1.1.0
* Ekafka版本: v1.0.6
* Emongo版本: v1.1.1

## 日志
EGO将日志分为两种类型，一种是系统日志，一种是业务日志。系统日志指的是`ego`框架或者`ego-component`组件。我们可以通过系统日志，就可以定位绝大部分问题。

我们可以根据以下字段，在日志系统里建立字段索引
| 字段名      | 用途        | 类型     | 
|----------|-----------|--------|
| ts       | 时间        |number|
| lv       | 日志级别      |string|
| msg      | 信息        |string|
| tid      | traceId   |string|
| comp     | 组件在框架中的名称 | string |
| compName | 组件在配置中的名称 | string | 
| code     | 错误码       | number|
| ucode    | 统一错误码     | number|
| cost     | 耗时        |number|
| method   | 方法        |string|
| error    | 错误信息      |string|
| addr     | 连接地址      |string|
| event    | 事件        |string|
|key| key       |string|
| ip       | ip        |string|
| name     | 名称        |string|
| peerIp   | 对端IP      | string|
| peerName | 对端名称      |string|
| type     | 类型        |string|


#### 1.1 第一步 Panic 日志观测
我们是不允许有`panic`日志，所以要对`panic`日志做好报警。
```bash
# 阿里云、腾讯云日志
lv:"panic"
# ClickVisual日志
`lv`='panic'
```

#### 1.2 第二步 Recover 日志观测
`EGO`在`HTTP`和`gRPC`服务有中间件，会对`panic`请求进行`recover`。
但由于`HTTP`服务如果暴露在公网，会出现`broken pipe`和`connection reset by peer`。所以这种情况下记录的`warn`级别。非这种情况记录的`error`级别。
因此`lv:"error" AND type:"recover"`这种日志是不允许存在的，需要做报警。
```bash
# 阿里云、腾讯云日志
lv:"error" AND type:"recover"
# ClickVisual日志
`lv`='error' AND `type`='recover'
```

#### 1.3 第三步统一错误码观测
我们会使用`ucode`这个字段，观测整个系统的错误码情况。
`ucode`是将`gRPC server`和`gRPC client`的`code`统一成和`HTTP code`一致。
我们只需要使用以下指令就能看到`HTTP Server`，`gRPC Server`（默认开启这些日志），和`HTTP Client`，`gRPC Client`（需要手动开启这些日志）
```bash
# 阿里云、腾讯云日志
ucode > 499
# ClickVisual日志
`ucode` > 499
```
如果发现问题
* 聚合`container name`看是哪个应用百分比最高
* 选择有问题的应用，聚合`method`，可以看是哪个方法存在问题
* 查看对应的`error`里面是否有错误信息提示
* 找到对应`tid`，将整个`tid`的日志查找出来，定位问题

#### 1.4 第四步慢日志观测
慢日志不代表系统有问题，但能说明系统的一些风险，也有可能有些慢日志是误报需要调整配置。
我们可以通过以下指令看到慢日志
```bash
# 阿里云、腾讯云日志
event:"slow"
# ClickVisual日志
`event`='slow' 
```
使用了慢日志的组件有
* HTTP服务: server.egin
* gRPC服务: server.egrpc
* HTTP客户端: client.ehttp
* gRPC客户端: client.egrpc
* MySQL组件: component.egorm
* Redis组件: component.eredis
* Mongo组件: component.emongo

我们使用以上的组件类型，可以定位不同慢日志情况。比如我们想要看HTTP服务慢日志，那么可以使用以下指令查询
```bash
# 阿里云、腾讯云日志
event:"slow" and comp:"server.egin"
# ClickVisual日志
`event`='slow' and `comp`='server.egin'
```

如果`MySQL`或者`Redis`等客户端出现了慢日志，可能有多种情况
* 指令在`MySQL`或者`Redis`运行慢了，那么需要优化指令，或者业务逻辑
* 指令在`MySQL`或者`Redis`运行不慢，但在客户端堵住了（可以通过metric监控指标确认），那么可能是连接池或者pod数不合理，需要调整连接池配置，或者增加pod数
* 指令在`MySQL`或者`Redis`运行不慢，客户端也没堵住，可能是`pod`的`CPU`不够。确认是高负载，还是`pod`的`CPU`给少了





