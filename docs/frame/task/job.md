# 短时任务Job
## 背景
通常我们有许多程序是短时任务，执行一下就结束。这种场景通常有以下两种方式：
* 执行某个一次性任务，例如：执行程序的安装，或者mock数据
* 将生命周期托管给例如k8s job或者xxljob，由他们控制job的执行时间

## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/examples/task/job)

## 用户代码
如果命令行参数里有 ``--job`` ，那么框架会优先执行这个 ``job``，停止所有的 ``server`` 和  ``cron`` 。 ``job`` 可以执行一个，也可以执行多个。执行一个方式 ``--job=jobname`` ，执行多个方式，用逗号分割 ``jobname``，例如： ``--job=jobname1,jobname2,jobname3``

```go
package main

import (
	"errors"
	"fmt"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/task/ejob"
	"go.uber.org/zap"
)

// export EGO_DEBUG=true && go run main.go --job=jobrunner
func main() {
	if err := ego.New().Job(NewJobRunner()).Run(); err != nil {
		elog.Error("start up", zap.Error(err))
	}
}

func NewJobRunner() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("jobrunner"),
		ejob.WithStartFunc(runner),
	)
}

func runner() error {
	fmt.Println("i am job runner")
	return errors.New("i am error")
}
``` 