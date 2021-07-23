# ETCD服务注册与发现原理

## 服务注册原理
在微服务架构下，主要有三种角色：服务提供者（RPC Server）、服务消费者（RPC Client）和服务注册中心（Registry），三者的交互关系请看下面这张图。

- RPC Server 提供服务，在启动时，根据服务的编译和配置信息，向 Registry 注册服务，并向 Registry 定期发送心跳汇报存活状态。
- RPC Client 调用服务，在启动时，根据配置文件的信息，向 Registry 订阅服务，把 Registry 返回的服务节点列表缓存在本地内存中，并与 RPC Sever 建立连接。
- 当 RPC Server 节点发生变更时，Registry 会同步变更，RPC Client 感知后会刷新本地内存中缓存的服务节点列表。
- RPC Client 从本地缓存的服务节点列表中，基于负载均衡算法选择一台 RPC Sever 发起调用。

![img_4.png](../../images/chapter3/img_4.png)

## 服务注册数据
| 名称 | 英文 | 示例 | 类型 |
| --- | --- | --- | --- |
| 环境 | env | dev | 环境变量 |
| 地区 | region | beijing | 环境变量 |
| 可用区 | zone | zone1 | 环境变量 |
| 地址 | ip | 192.168.1.1 | 环境变量 |
| 端口 | port | 8080 | 环境变量 |
| 协议 | scheme | gRPC | 配置变量 |
| 权重 | weight | 100 | 配置变量 |
| 部署组 | deployment | red | 配置变量 |
| 框架版本号 | frameVersion | 2.0 | 编译变量 |
| 应用版本号 | appVersion | ahkfasgasdf | 编译变量 |
| 编译时间 | buildTime | 2020-10-14 10:00:00 | 编译变量 |
| 启动时间 | startTime | 2020-10-14 11:00:00 | 编译变量 |

通过这些信息，我们能够很方便的获得Provider的基本情况，并对他进行改变。例如可以调节流量比例、调度流量区域、灰度版本、排查应用的版本、编译时间、启动时间。

## 启动一个服务例子
example: 
### 编译


```bash
~ etcdctl get "/ego" --prefix
 /ego/hello/providers/grpc://127.0.0.1:9003
{
    "name":"hello",
    "scheme":"grpc",
    "address":"127.0.0.1:9003",
    "weight":100,
    "enable":true,
    "healthy":true,
    "region":"huabei",
    "zone":"ali-3",
    "kind":1,
    "deployment":"",
    "group":"",
    "metadata":{
        "appHost":"127.0.0.1",
        "appMode":"dev",
        "appVersion":"5c569c3a7427d3266cfd15f23b37f924f083b785-dirty",
        "buildTime":"2021-06-23 22:58:43",
        "egoVersion":"v0.5.8",
        "startTime":"2021-06-23 22:58:54"
    },
}
```

