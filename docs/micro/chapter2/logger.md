## 如何玩好日志

* 采集日志
* 日志规范
* 日志收敛
* 日志trace

### 采集日志

DockerEngine 本身具有 LogDriver 功能，可通过配置不同的 LogDriver 将容器的 stdout 通过 DockerEngine 写入到远端存储，以此达到日志采集的目的。

* 这种方式的可定制化、灵活性、资源隔离性都很低，一般不建议在生产环境中使用；
* 业务直写是在应用中集成日志采集的 SDK，通过 SDK 直接将日志发送到服务端。这种方式省去了落盘采集的逻辑，也不需要额外部署 Agent，对于系统的资源消耗最低，但由于业务和日志 SDK
  强绑定，整体灵活性很低，一般只有日志量极大的场景中使用；
* DaemonSet 方式在每个 node 节点上只运行一个日志 agent，采集这个节点上所有的日志。DaemonSet 相对资源占用要小很多，但扩展性、租户隔离性受限，比较适用于功能单一或业务不是很多的集群；
* Sidecar 方式为每个 POD 单独部署日志 agent，这个 agent 只负责一个业务应用的日志采集。Sidecar 相对资源占用较多，但灵活性以及多租户隔离性较强，建议大型的 K8s 集群或作为 PaaS
  平台为多个业务方服务的集群使用该方式。

总结下来：

* DockerEngine 直写一般不推荐；
* 业务直写推荐在日志量极大的场景中使用；
* DaemonSet 一般在中小型集群中使用；
* Sidecar 推荐在超大型的集群中使用。

阿里云K8S采集日志方式如下所示：

![img.png](../../images/collectlogger.png)

### EGO的日志说明

* 推荐采用DaemonSet方式
* 支持三种日志输出方式
    * 文件（默认）
    * 终端（需要改日志输出方式）
    * 直接写入到数据源
* 日志区分
    * ego.sys      (框架日志)
    * default.log （业务日志）


### 日志输出：Stdout or 文件
和虚拟机/物理机不同，K8s 的容器提供标准输出和文件两种方式。在容器中，标准输出将日志直接输出到 stdout 或 stderr，而 DockerEngine 接管 stdout 和 stderr 文件描述符，将日志接收后按照 DockerEngine 配置的 LogDriver 规则进行处理；日志打印到文件的方式和虚拟机/物理机基本类似，只是日志可以使用不同的存储方式，例如默认存储、EmptyDir、HostVolume、NFS 等。
虽然使用 Stdout 打印日志是 Docker 官方推荐的方式，但大家需要注意：这个推荐是基于容器只作为简单应用的场景，实际的业务场景中我们还是建议大家尽可能使用文件的方式，主要的原因有以下几点：

* Stdout 性能问题，从应用输出 stdout 到服务端，中间会经过好几个流程（例如普遍使用的 JSON LogDriver）：应用 stdout -> DockerEngine -> LogDriver -> 序列化成 JSON -> 保存到文件 -> Agent 采集文件 -> 解析 JSON -> 上传服务端。整个流程相比文件的额外开销要多很多，在压测时，每秒 10 万行日志输出就会额外占用 DockerEngine 1 个 CPU 核；

* Stdout 不支持分类，即所有的输出都混在一个流中，无法像文件一样分类输出，通常一个应用中有 AccessLog、ErrorLog、InterfaceLog（调用外部接口的日志）、TraceLog 等，而这些日志的格式、用途不一，如果混在同一个流中将很难采集和分析；

* Stdout 只支持容器的主程序输出，如果是 daemon/fork 方式运行的程序将无法使用 stdout；

* 文件的 Dump 方式支持各种策略，例如同步/异步写入、缓存大小、文件轮转策略、压缩策略、清除策略等，相对更加灵活。

因此我们建议线上应用使用文件的方式输出日志，Stdout 只在功能单一的应用或一些 K8s 系统/运维组件中使用。

#### 日志等级
日志等级是用来区分日志对应事件严重程度的说明，这是所有日志中必须具备的一个选项。EGO采用的是zap的日志等级方式：
* Debug日志：用于线下调试的日志信息，用于分析应用执行逻辑，线上应用切勿开启；
* Info日志：应用执行过程中的详细信息，一般通过该信息可以看到每个请求的主要执行过程；
* Warn日志：潜在的危险或值得关注的信息（比较核心的路径）；
* Error日志：非预期中的错误，此种错误可能导致部分系统异常但不会影响核心业务和系统正常运行；
* DPanic日志:
* Panic日志: 此种错误可能导致部分系统异常，可能会影响核心业务和系统正常运行，遇到此种错误应当立即报警并人工介入处理；
* Fatal日志：用来输出非常严重或预期中不会发生的错误，遇到此种错误应当立即报警并人工介入处理；


### refer
* [6 个 K8s 日志系统建设中的典型问题，你遇到过几个？](https://developer.aliyun.com/article/718735)
* [一文看懂 K8s 日志系统设计和实践](https://developer.aliyun.com/article/727594)
* [9 个技巧，解决 K8s 中的日志输出问题](https://developer.aliyun.com/article/747821)
* [详解 K8s 日志采集最佳实践](https://blog.csdn.net/JKX_geek/article/details/104858769)
* [再次升级！阿里云Kubernetes日志解决方案](https://blog.csdn.net/maoreyou/article/details/80487138)
* https://reading.developerlearning.cn/reading/29-2019-01-23-opentracing-jaeger-in-go/
* http://bigbully.github.io/Dapper-translation/
* https://blog.csdn.net/yunqiinsight/article/details/80134045
* https://www.jianshu.com/p/752c17eb33f3
* https://www.jaegertracing.io/docs/1.14/sampling/

