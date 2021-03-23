# 定时任务 Cron

## 1 Example

[Example](https://github.com/gotomicro/ego/tree/master/examples/task/cron)

ego 版本：`ego@v0.4.1`

## 2 定时任务配置

```go
type Config struct {
	// Required. 触发时间
	//	默认最小单位为分钟.比如:
	//		"* * * * * *" 代表每分钟执行
	//	如果 EnableSeconds = true. 那么最小单位为秒. 示例:
	//		"*/3 * * * * * *" 代表每三秒钟执行一次
	Spec string

	WaitLockTime   time.Duration // 抢锁等待时间，默认 4s
	LockTTL        time.Duration // 租期，默认 16s
	RefreshGap     time.Duration // 锁刷新间隔时间， 默认 4s
	WaitUnlockTime time.Duration // 解锁等待时间，默认 1s

	// 任务时间重叠策略, 可选项:
	// 		skip，queue，concurrent
	// 如果上一个任务执行较慢，到达了新任务执行时间，那么新任务选择跳过，排队，并发执行的策略，新任务默认选择skip策略
	DelayExecType         string
	// 是否分布式任务，默认否
	// 如果设置为 true. 那么需要设置 ecron.WithLock Option
	// 框架会使用分布式锁保证同时只有一个节点在执行当前分布式任务
	EnableDistributedTask bool
	EnableImmediatelyRun  bool   // 是否立刻执行，默认否
	EnableSeconds         bool   // 是否使用秒作解析器，默认否
}
```

## 3 常规定时任务

### 3.1 用户配置

```toml
[cron.test]
enableDistributedTask = false			# 是否分布式任务，默认否，如果存在分布式任务，会只执行该定时人物
enableImmediatelyRun = false			# 是否立刻执行，默认否
enableSeconds = false				# 是否使用秒作解析器，默认否
spec = "*/5 * * * * *"					# 执行时间
delayExecType = "skip"					# skip，queue，concurrent，如果上一个任务执行较慢，到达了新任务执行时间，那么新任务选择跳过，排队，并发执行的策略，新任务默认选择skip策略
```

### 3.2 用户代码

配置创建一个 `{{你的配置key}}` 的配置项，其中内容按照上文 HTTP 的配置进行填写。以上这个示例里这个配置 key 是`cron.test`

代码中创建一个 `cron` 服务， ecron.Load("{{你的配置key}}").Build() ，代码中的 `key` 和配置中的 `key` 。创建完 `cron` 后， 将他添加到 `ego new` 出来应用的 `Schedule` 方法中。

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/gotomicro/ego/task/ecron"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New().Cron(cronJob1(), cronJob2()).Run()
	if err != nil {
		elog.Panic("startup engine", elog.FieldErr(err))
	}
}

// 异常任务
func cronJob1() ecron.Ecron {
	job := func(ctx context.Context) error {
		elog.Info("info job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		elog.Warn("warn job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		fmt.Println("run job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		return errors.New("exec job1 error")
	}

	cron := ecron.Load("cron.test").Build(ecron.WithJob(job))
	return cron
}

// 正常任务
func cronJob2() ecron.Ecron {
	job := func(ctx context.Context) error {
		elog.Info("info job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		elog.Warn("warn job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		fmt.Println("run job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		return nil
	}

	cron := ecron.Load("cron.test").Build(ecron.WithJob(job))
	return cron
}
```

## 4 分布式定时任务

### 4.1 用户配置

```toml
[cron.test]
enableDistributedTask = true          # 是否分布式任务，默认否，如果存在分布式任务，会只执行该定时人物
enableImmediatelyRun = false        # 是否立刻执行，默认否
delayExecType = "skip"  # skip，queue，concurrent，如果上一个任务执行较慢，到达了新任务执行时间，那么新任务选择跳过，排队，并发执行的策略，新任务默认选择skip策略
enableSeconds = true # 启用秒单位
spec = "*/3 * * * * *"

[redis.test]
debug = true # ego增加redis debug，打开后可以看到，配置名、地址、耗时、请求数据、响应数据
addr = "127.0.0.1:6379"
enableAccessInterceptor = true
enableAccessInterceptorReq = true
enableAccessInterceptorRes = true
```

### 4.2 用户代码
配置创建一个 ``{{你的配置key}}`` 的配置项，其中内容按照上文HTTP的配置进行填写。以上这个示例里这个配置key是``cron.test``

代码中创建一个 ``cron`` 服务， ecron.Load("{{你的配置key}}").Build() ，代码中的 ``key`` 和配置中的 ``key`` 。创建完 ``cron`` 后， 将他添加到 ``ego new`` 出来应用的 ``Schedule`` 方法中。

```go
package main

import (
	"context"
	"log"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/task/ecron"

	"github.com/gotomicro/ego-component/eredis"
	"github.com/gotomicro/ego-component/eredis/ecronlock"
)

var (
	redis  *eredis.Component
	locker *ecronlock.Component
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	err := ego.New().Invoker(initRedis).Cron(cronJob()).Run()
	if err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func initRedis() error {
	redis = eredis.Load("redis.test").Build()
	// 构造分布式任务锁，目前已实现redis版本. 如果希望自定义，可以实现 ecron.Lock 接口
	locker = ecronlock.DefaultContainer().Build(ecronlock.WithClient(redis))
	return nil
}

func cronJob() ecron.Ecron {
	cron := ecron.Load("cron.test").Build(
		// 设置分布式锁
		ecron.WithLock(locker.NewLock("ego-component:cronjob:syncXxx")),
		ecron.WithJob(helloWorld),
	)
	return cron
}

func helloWorld(ctx context.Context) error {
	log.Println("cron job running")
	return nil
}
```

<Vssue title="Task-cron" />
