# 定时任务Cron
## 1 Example
[项目地址](https://github.com/gotomicro/ego/tree/master/examples/task/cron)
ego版本：``ego@v0.3.11``

## 2 定时任务配置
```go
type Config struct {
	WaitLockTime          time.Duration // 抢锁等待时间，默认60s
	LockTTL               time.Duration // 租期，默认60s
	LockDir               string        // 定时任务锁目录
	RefreshTTL            time.Duration // 刷新ttl，默认60s
	WaitUnlockTime        time.Duration // 抢锁等待时间，默认1s
	DelayExecType         string        // skip，queue，concurrent，如果上一个任务执行较慢，到达了新的任务执行时间，那么新的任务选择跳过，排队，并发执行的策略
	EnableDistributedTask bool          // 是否分布式任务，默认否，如果存在分布式任务，会只执行该定时人物
	EnableImmediatelyRun  bool          // 是否立刻执行，默认否
	EnableWithSeconds     bool          // 是否使用秒作解析器，默认否
}
```
## 3 常规定时任务
### 3.1 用户配置
```toml
[cron.test]
enableDistributedTask = false          # 是否分布式任务，默认否，如果存在分布式任务，会只执行该定时人物
enableImmediatelyRun = false        # 是否立刻执行，默认否
enableWithSeconds = false      # 是否使用秒作解析器，默认否
delayExecType = "skip"  # skip，queue，concurrent，如果上一个任务执行较慢，到达了新任务执行时间，那么新任务选择跳过，排队，并发执行的策略，新任务默认选择skip策略
```

### 3.2 用户代码
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

## 4 分布式定时任务
### 4.1 用户配置
```toml
[cron.test]
enableDistributedTask = true          # 是否分布式任务，默认否，如果存在分布式任务，会只执行该定时人物
enableImmediatelyRun = false        # 是否立刻执行，默认否
enableWithSeconds = false      # 是否使用秒作解析器，默认否
delayExecType = "skip"  # skip，queue，concurrent，如果上一个任务执行较慢，到达了新任务执行时间，那么新任务选择跳过，排队，并发执行的策略，新任务默认选择skip策略
```

### 4.2 用户代码
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
	lock := ecronlock.Load("").Build(ecronlock.WithClientRedis(invoker.Redis))
	cron := ecron.Load("cron.test").Build(ecron.WithLocker(lock))
	cron.Schedule(ecron.Every(time.Second*10), ecron.FuncJob(execJob))
	return cron
}

func execJob() error {
	elog.Info("info job")
	elog.Warn("warn job")
	fmt.Println("run job")
	return errors.New("exec job1 error")
}
```


<Vssue title="Task-cron" />