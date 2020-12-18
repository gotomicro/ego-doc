# 定时任务Cron
## Example
[项目地址](https://github.com/gotomicro/ego/tree/master/examples/task/cron)

## 定时任务配置
```go
type Config struct {
	WithSeconds      bool          // 是否使用秒作解析器，默认否
	ConcurrentDelay  int           // 并发延迟，默认是执行超过定时时间后，下次执行的任务会跳过
	ImmediatelyRun   bool          // 是否立刻执行，默认否
	DistributedTask  bool          // 是否分布式任务，默认否，如果存在分布式任务，则会解析嵌入的etcd配置
	WaitLockTime     time.Duration // 抢锁等待时间，默认0s
	Endpoints        []string      // etcd地址
	ConnectTimeout   time.Duration // 连接超时时间，默认5s
	Secure           bool          // 是否安全通信，默认false
	AutoSyncInterval time.Duration // 自动同步member list的间隔
	TTL              int           // 过期时间，单位：s，默认失效时间为0s
	WorkerLockDir    string        // 定时任务锁目录
}
```

## 用户配置
```toml
[cron.test]
    withSeconds = false
    concurrentDelay= 1
    immediatelyRun = false
    distributedTask = false
    waitLockTime = "1s"
```

## 用户代码
配置创建一个 ``{{你的配置key}}`` 的配置项，其中内容按照上文HTTP的配置进行填写。以上这个示例里这个配置key是``cron.test``

代码中创建一个 ``cron`` 服务， ecron.Load("{{你的配置key}}").Build() ，代码中的 ``key`` 和配置中的 ``key`` 。创建完 ``cron`` 后， 将他添加到 ``ego new`` 出来应用的 ``Schedule`` 方法中。

```go
package main

import (
	"errors"
	"fmt"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/task/ecron"
	"time"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New().Cron(cron1()).Run()
	if err != nil {
		elog.Panic("startup engine", elog.Any("err", err))
	}
}

func cron1() ecron.Ecron {
	cron := ecron.Load("cron.test").Build()
	cron.Schedule(ecron.Every(time.Second*10), ecron.FuncJob(execJob))
	cron.Schedule(ecron.Every(time.Second*10), ecron.FuncJob(execJob2))
	return cron
}

// 异常任务
func execJob() error {
	elog.Info("info job")
	elog.Warn("warn job")
	fmt.Println("run job")
	return errors.New("exec job1 error")
}

// 正常任务
func execJob2() error {
	elog.Info("info job2")
	elog.Warn("warn job2")
	fmt.Println("run job2")
	return nil
}
```

<Vssue title="Task-cron" />