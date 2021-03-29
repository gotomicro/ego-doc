# 1 编译
## 1.1 Go编译注入信息
Go微服务的编译是微服务的第一步，也是比较重要的一个环节。我们可以在编译的时候注入很多编译信息，例如应用名称、应用版本号、框架版本号、编译所在机器、编译时间，我们可以直接注入到二进制里。编译完成后，我们可以使用`./micro --version` ，查看该服务的基本情况，如下图所示。

![image](../../images/buildversion.png)

我们还可以在微服务启动后，将这些编译信息写入prometheus或者etcd中。当线上出现什么问题的时候，我们能够快速知道微服务在线上使用的哪个版本、编译在什么时间，提升我们排查微服务问题的速度。

接下来我们就来看下如何在Go微服务里编译这些信息

## 1.2 常用编译指令
`go build`指令比较多。我们把微服务里常用的命令展示在下表：
| 参数 | 备注 |
| --- | --- |
| -o | 目标地址 |
| -race | 开启竞态检测 |
| -ldflags | 传递参数 |
| -n | 打印编译时会用到的所有命令，但不真正执行 |
| -x | 打印编译时会用到的所有命令 |
| -tag | 根据tag版本编译 |

### -o
编译到指定地址
```bash
go build -o micro
```
### -race
开启竞态检查编译。通过这个编译方式。你的程序可以在运行的时候崩溃
```bash
go build -o micro -race
curl http://127.0.0.1:8080/race
```
我们开启race编译后，访问该地址，就可以看到代码中出现race的报错

### -ldflags
- -w 去掉DWARF调试信息，得到的程序就不能用gdb调试了
-  -s 去掉符号表,panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果
- -X 设置包中的变量值
```bash
 go build -o micro -ldflags "-X main.buildName=micro\
 -X main.buildGitRevision=f8c315083e7b739f0f055ee46a747c8e109d7539-dirty\
 -X main.buildStatus=Modified -X main.buildUser=`whoami` \
 -X main.buildHost=`hostname -f` -X main.buildTime=`date +%Y-%m-%d--%T`"
```

### 编译演示代码
```go
package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"runtime"
)

var (
	buildName        = "unknown"
	buildGitRevision = "unknown"
	buildUser        = "unknown"
	buildHost        = "unknown"
	buildStatus      = "unknown"
	buildTime        = "unknown"
)

var (
	version bool
	run     bool
)

func init() {
	pflag.BoolVarP(&version, "version", "v", false, `查看版本号`)
	pflag.BoolVarP(&run, "run", "r", false, `运行程序`)
	pflag.Parse()
}

func main() {
	if version == true {
		fmt.Println(LongForm())
	}
	if run == true {
		fmt.Println("go to micro")
	}
}

func LongForm() string {
	return fmt.Sprintf(`Name: %v
GitRevision: %v
User: %v@%v
GolangVersion: %v
BuildStatus: %v
BuildTime: %v
`,
		buildName,
		buildGitRevision,
		buildUser,
		buildHost,
		runtime.Version(),
		buildStatus,
		buildTime)
}
```

### -tag
用于编译打tag，灰度测试代码使用。例如
```bash
go build -o micro -tag="build1"
```

## 1.3 EGO编译指令
### 脚本
[脚本代码路径](https://github.com/gotomicro/ego/blob/master/scripts/build/report_build_info.sh)

脚本核心代码如下：
```bash
echo "github.com/gotomicro/ego/core/eapp.appName=${APP_NAME}"
echo "github.com/gotomicro/ego/core/eapp.buildVersion=${VERSION}"
echo "github.com/gotomicro/ego/core/eapp.buildAppVersion=${BUILD_GIT_REVISION}"
echo "github.com/gotomicro/ego/core/eapp.buildStatus=${tree_status}"
echo "github.com/gotomicro/ego/core/eapp.buildUser=$(whoami)"
echo "github.com/gotomicro/ego/core/eapp.buildHost=$(hostname -f)"
echo "github.com/gotomicro/ego/core/eapp.buildTime=$(date '+%Y-%m-%d--%T')"
```

### 获取框架版本代码
应用大部分信息是通过编译时期通过脚本获取信息注入到二进制包的变量里，但应用获取``ego``框架版本则是通过遍历``runtime/debug``包的版本信息，匹配后获取对应的``ego``版本
```go
egoVersion := "unknown version"
info, ok := debug.ReadBuildInfo()
if ok {
    for _, value := range info.Deps {
        if value.Path == "github.com/gotomicro/ego" {
            egoVersion = value.Version
        }
    }
}
```

### 信息展示
最终我们会将应用的编译信息和环境变量信息写入到以下prometheus中，方便我们查询业务的基本情况。
```go
GaugeVecOpts{
    Namespace: DefaultNamespace,
    Name:      "build_info",
    Labels:    []string{"name", "mode", "region", "zone", "app_version", "ego_version", "start_time", "build_time", "go_version"},
}.Build()
```