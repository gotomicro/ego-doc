# 配置
## Example
* 动态配置
    * [获取单个配置信息](https://github.com/gotomicro/ego/tree/master/examples/config/oneline-by-file-watch) 
    * [获取结构体信息](https://github.com/gotomicro/ego/tree/master/examples/config/struct-by-file-watch)
* 静态配置
    * [获取单个配置信息](https://github.com/gotomicro/ego/tree/master/examples/config/oneline-by-file)
    * [获取结构体信息](https://github.com/gotomicro/ego/tree/master/examples/config/struct-by-file)
    
## 系统内置
* 文本配置
* 框架默认开启动态配置
* 框架默认内志支持文本更改日志级别，动态生效

## 代码示例
### 动态配置获取单个配置信息
```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
	"time"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New(ego.WithHang(true)).Invoker(func() error {
		go func() {
			// 循环打印配置
			for {
				time.Sleep(3 * time.Second)
				peopleName := econf.GetString("people.name")
				elog.Info("people info", elog.String("name", peopleName), elog.String("type", "onelineByFileWatch"))
			}
		}()
		return nil
	}).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
```

### 动态配置获取结构体配置信息
```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/elog"
	"time"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New(ego.WithHang(true)).Invoker(func() error {
		p := People{}
		// 初始化
		err := econf.UnmarshalKey("people", &p)
		if err != nil {
			panic(err.Error())
		}
		// 监听
		econf.OnChange(func(config *econf.Configuration) {
			err := config.UnmarshalKey("people", &p)
			if err != nil {
				elog.Panic("unmarshal", elog.FieldErr(err))
			}
		})

		go func() {
			// 循环打印配置
			for {
				time.Sleep(1 * time.Second)
				elog.Info("people info", elog.String("name", p.Name), elog.String("type", "structByFileWatch"))
			}
		}()
		return nil
	}).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

type People struct {
	Name string
}
```

<Vssue title="Config" />
