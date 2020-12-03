# 编译
## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/example/build)

使用EGO框架的应用会在编译期注入许多必要信息，方便后续排查问题。该方案被大量Go应用所使用，例如istio、prometheus等。我们使用的编译脚本核心内容如下所示。

## 编译脚本
```bash
go build -o bin/hello -pkgdir=/Users/askuy/go/pkg/linux_amd64 -ldflags -extldflags -static  -X github.com/gotomicro/ego/core/app.appName=hello -X github.com/gotomicro/ego/core/app.buildVersion=925d5b27ff35b4490494ba78ceb897e02cb12d92-dirty -X github.com/gotomicro/ego/core/app.buildAppVersion=925d5b27ff35b4490494ba78ceb897e02cb12d92-dirty -X github.com/gotomicro/ego/core/app.buildStatus=Modified -X github.com/gotomicro/ego/core/app.buildTag= -X github.com/gotomicro/ego/core/app.buildUser=askuy -X github.com/gotomicro/ego/core/app.buildHost=askuydeMacBook-Pro.local -X github.com/gotomicro/ego/core/app.buildTime=2020-12-03--17:26:24
```


## 查看编译版本信息
![](../../images/buildversion.png)

## 查看运行时信息
