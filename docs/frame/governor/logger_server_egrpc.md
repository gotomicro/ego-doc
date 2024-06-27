# Server gRPC 日志

## 字段信息
| 字段名      | 用途        | 类型     | 例子                               |
|----------|-----------|--------|----------------------------------|
| ts       | 时间        |number| 1719372895                       |
| lv       | 日志级别      |string| info                             |
| msg      | 信息        |string| access                           |
| tid      | traceId   |string| edff12dca2cd72bb710a1c1d4e5e8cf4 |
| comp     | 组件在框架中的名称 | string | server.egrpc                     |
| compName | 组件在配置中的名称 | string | server.grpc                      |
| code     | 错误码       | number| 5                                |
| ucode    | 统一错误码     | number| 500                              |
| cost     | 耗时        |number| 1.23                             |
| method   | 方法        |string| GET./api/file/:fileGuid          |
| error    | 错误信息      |string|                                  |
| addr     | 连接地址      |string| /api/file/abcd?test=123          |
| event    | 事件        |string| normal,slow                      |
|type|类型|string|recover|
| ip       | ip        |string| 1.1.1.1                          |
| name     | 名称        |string|                                  |
| peerIp   |对端IP| string| 10.1.1.1                         |
| peerName |对端名称|string| svc-user                         |
| key      |类型|string| GET,POST                         |
|req| 请求数据|any||
|res|响应数据|any||

## Panic日志
```bash
comp:"server.egrpc" and lv:"error" and type:"recover"
```
出现这种日志，说明了`gRPC`服务出现了`panic`，然后被`recover`，这种错误需要报警。
我们需要观察这个错误的`stack`字段。

## 5xx日志
```bash
comp:"server.egrpc" and ucode > 499
```
出现这种日志，说明了`HTTP`服务出现了`5xx`，这种错误需要报警。
* 聚合`container name`看是哪个应用百分比最高
* 选择有问题的应用，聚合`method`，可以看是哪个方法存在问题
* 查看对应的`error`里面是否有错误信息提示
* 找到对应`tid`，将整个`tid`的日志查找出来，定位问题


## 慢日志
```bash
comp:"server.egrpc" and event:"slow"
```
出现这种日志，说明了`HTTP`服务出现了`slow`日志，这种慢日志需要报警。
我们需要观察这个耗时的`cost`字段。
* 聚合`container name`看是哪个应用百分比最高
* 选择有问题的应用，聚合`method`，可以看是哪个耗时

