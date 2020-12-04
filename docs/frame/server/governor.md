## 治理服务
## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/example/server/governor)

## 背景
``Go``不像``Java``和``PHP``，有虚拟机帮助程序员对程序运行的内部情况的观测，但这种观测对于程序员而言排查故障，解决性能问题是非常重要的。

``EGO``着眼于可观测性，在各个组件里引入拦截器，提取有用信息，通过实现一个治理服务，将程序运行数据吐出来，方便用户做治理平台，排查各类问题。

## 可观测数据
|路由| 说明|
| --- | --- |
| / | 展示所有可观测的路由 |
| /metrics | 监控数据 |
| /debug/pprof/* | pprorf信息 |
| /config/json | config json数据 |
| /config/raw | config 原始数据 |
| /module/info | 应用依赖模块信息 |
| /build/info | 应用编译信息 |
| /env/info | 应用环境信息 |
| /code/info | 状态码信息，待完成 |
| /component/info | 组件信息，待完成 |

## 用户配置
```
[server.governor]
  host = "0.0.0.0"
  port = 9003
```


## 用户代码
配置创建一个 ``{{你的配置key}}`` 的配置项，其中内容按照上文HTTP的配置进行填写。以上这个示例里这个配置key是``server.governor``

代码中创建一个 ``governor`` 服务， egin.Load("{{你的配置key}}").Build() ，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``http`` 服务后， 将他添加到 ``ego new`` 出来应用的 ``Serve`` 方法中，之后使用的方法和 ``gRPC`` 就完全一致。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
	"github.com/gotomicro/ego/server/egovernor"
)

func main() {
	if err := ego.New().
		Serve(
			egovernor.Load("server.governor").Build(),
			serverHttp(),
		).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func serverHttp() *egin.Component {
	server := egin.Load("server.http").Build()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
		return
	})
	return server
}
```

## 查看运行的metric信息
我们启动服务

![](../../images/buildrun.png)

请求治理端口的/metrics接口，可以看到 ``ego_build_info`` 的信息，这里会将编译信息放入到 ``prometheus`` 中，并且还会把运行时的环境信息和启动时间也加入进来。
```
# HELP ego_build_info 
# TYPE ego_build_info gauge
ego_build_info{app_version="b0807b91aca95b6eb6daafa9195c467fac0c350b-dirty",build_time="2020-12-04 11:48:35",ego_version="0.1.0",go_version="go1.15.2",mode="dev",name="hello",region="huabei",start_time="2020-12-04 11:49:02",zone="ali-3"} 1.607053742679e+12
```