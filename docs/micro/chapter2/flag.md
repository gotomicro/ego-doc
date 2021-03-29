# 2.1 启动参数
## 2.1.1 介绍
启动参数主要用于应用启动的必要参数，或者根据不同环境选择的的不同参数。
典型的场景例如：
* --help 查看帮助文档
* --version 查看应用版本
* --config 查看应用的启动配置（必要参数，根据环境变化的参数）
* --host 选择应用的主机IP（根据环境变化的参数）

### 2.1.2 最佳实践--统一启动参数
因为命令行参数通常是研发人员来指定，但每个项目都指定大量的启动参数，是一个很繁琐的事情。

我们通常会将公用的配置，例如环境、Region、Zone、配置路径、启动IP等信息，通过基础设施，统一设置成环境变量，精简研发人员的启动参数。

同时我们可以通过环境变量设置一个公司的研发规范。例如配置在``dev``环境使用``dev.toml``，``prod``环境使用``prod.toml``。

### 2.1.3 最佳实践--常用微服务启动参数
这里列出了常用微服务的启动参数

|命令行参数|环境变量|默认参数| 描述 |
| --- | --- | --- |--- |
|config|EGO_CONFIG_PATH|config/local.toml| 配置路径|
|host|EGO_HOST|0.0.0.0| 启动IP|
|watch|EGO_WATCH|true| 默认监听|
|debug|EGO_DEBUG|false| 是否开启调试模式 |
|ego_name|EGO_NAME|filepath.Base(os.Args[0])| 应用名|
|ego_mode|EGO_MODE|空| 环境 |
|ego_region|EGO_REGION|空| 地区 |
|ego_zone|EGO_ZONE|空| 可用区 |
|ego_log_path|EGO_LOG_PATH|./logs| 配置路径 |
|ego_log_add_app|EGO_LOG_ADD_APP|false| 日志里是否添加应用名 |
|ego_trace_id_name|EGO_TRACE_ID_NAME|x-trace-id| 链路名称 |

## 2.1.4 优先级
优先级: 命令行参数 > 环境变量 > 默认参数





