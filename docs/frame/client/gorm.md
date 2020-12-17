# GORM
## Example
[项目地址](https://github.com/gotomicro/ego-component/tree/master/egorm/examples/gorm)

## GORM配置
```go
type Config struct {
	Dialect                      string        // 选择数据库种类，默认mysql
	DSN                          string        // DSN地址: mysql://root:secret@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s
	Debug                        bool          // Debug开关，默认关闭
	RawDebug                     bool          // 原生Debug开关，默认关闭
	MaxIdleConns                 int           // 最大空闲连接数，默认10
	MaxOpenConns                 int           // 最大活动连接数，默认100
	ConnMaxLifetime              time.Duration // 连接的最大存活时间，默认300s
	OnFail                       string        // 创建连接的错误级别，=panic时，如果创建失败，立即panic，默认连接不上panic
	SlowLogThreshold             time.Duration // 慢日志阈值，默认500ms
	DisableMetricInterceptor     bool          // 关闭指标采集，默人false，也就是说默认采集监控
	DisableTrace                 bool          // 关闭链路追踪，默认false，也就是说默认采集trace
	EnableDetailSQL              bool          // 记录错误sql时,是否打印包含参数的完整sql语句，select * from aid = ?;
	EnableAccessInterceptor      bool          // 是否开启，记录请求数据
	EnableAccessInterceptorReply bool          // 是否开启记录响应参数
	EnableAccessInterceptorReq   bool          // 是否开启记录请求参数
}
```


## 用户配置
```toml
[mysql.test]
   debug = true # ego重写gorm debug，打开后可以看到，配置名、地址、耗时、请求数据、响应数据
   dsn = "root:root@tcp(127.0.0.1:3306)/ego?charset=utf8&parseTime=True&loc=Local&readTimeout=1s&timeout=1s&writeTimeout=3s"
```

## 用户代码
::: tip
客户端组件均使用go mod子包管理，使用该组件，一定要使用下面的go get命令
:::

```bash
go get github.com/gotomicro/ego-component/egorm
```

配置创建一个 ``{{你的配置key}}`` 的配置项，其中内容按照上文gorm的配置进行填写。

代码中创建一个 ``gorm`` 实例 ``egorm.Load("{{你的配置key}}").Build()``，代码中的 ``key`` 和配置中的 ``key`` 要保持一致。创建完 ``gorm`` 实例后，就可以直接使用他对 ``db`` 进行 ``crud`` 。


```go
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
)

/**
1.新建一个数据库叫test
2.执行以下example，export EGO_DEBUG=true && go run main.go --config=config.toml
*/
type User struct {
	Id       int    `gorm:"not null" json:"id"`
	Nickname string `gorm:"not null" json:"name"`
}

func (User) TableName() string {
	return "user2"
}

func main() {
	err := ego.New().Invoker(
		openDB,
		testDB,
	).Run()
	if err != nil {
		elog.Error("startup", elog.Any("err", err))
	}
}

var gormDB *egorm.Component

func openDB() error {
	gormDB = egorm.Load("mysql.test").Build()
	models := []interface{}{
		&User{},
	}
	gormDB.SingularTable(true)
	gormDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	gormDB.Create(&User{
		Nickname: "ego",
	})
	return nil
}

func testDB() error {
	var user User
	err := gormDB.Where("id = ?", 100).Find(&user).Error
	elog.Info("user info", elog.String("name", user.Nickname))
	return err
}
```