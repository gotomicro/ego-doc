## 1 架构演进
互联网的WEB架构演进可以分为三个阶段：单体应用时期、垂直应用时期、微服务时期。


单体应用时期一般处于一个公司的创业初期，他的好处就是运维简单、开发快速、能够快速适应业务需求变化。但是当业务发展到一定程度后，会发现许多业务会存在一些莫名奇妙的耦合，例如你修改了一个支付模块的函数，结果登录功能挂了。为了避免这种耦合，会将一些功能模块做一个垂直拆分，进行业务隔离，彼此之间功能相互不影响。但是在业务发展过程中，会发现垂直应用架构有许多相同的功能，需要重复开发或者复制粘贴代码。所以要解决以上复用功能的问题，我们可以将同一个业务领域内功能抽出来作为一个单独的服务，~~​~~服务之间使用RPC进行远程调用，这就是我们常所说的微服务架构。

总的来说，我们可以将这三个阶段总结为以下几点。单体应用架构快速、简单，但~~​~~耦合性强；垂直应用架构隔离性、稳定性好，但复制粘贴代码会比较多；微服务架构可以说是兼顾了垂直应用架构的隔离性、稳定性，并且有很强的复用性能力。可以说微服务架构是公司发展壮大后，演进到某种阶段的必然趋势。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631804426745-bfdece05-88ce-4a80-837c-ea3ea76988fd.png#clientId=ue5f57623-9a54-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=319&id=u99f16720&margin=%5Bobject%20Object%5D&name=image.png&originHeight=319&originWidth=1178&originalType=binary&ratio=1&rotation=0&showTitle=false&size=36623&status=done&style=none&taskId=u67863f76-ed23-4e1b-8f51-37195047aae&title=&width=1178)


但微服务真的那么美好吗？我们可以看到一个单体架构和微服务架构的对比图。在左图我们可以看到一个业务可以通过Nginx+服务器+数据库就能实现业务需求。但是在右图微服务架构中，我们完成一个业务需要引入大量的组件，比如在中间这一块我们会引入DNS、HPA、ConfigMap等、下面部分引入了存储组件Redis、MySQL、Mongo等。以前~~​~~单体应用时期我们可能直接上机器看日志或上机器上查看资源负载监控，但是到了微服务阶段，应用太多了，肯定不能这么去操作，这个时候我们就需要引入ELK、Prometheus、Grafana、Jaeger等各种基础设施，来更方便地对我们的服务进行观测。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631261406902-fbfba61d-5186-4efa-999e-6ab1a638e2ff.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=514&id=u117c8c43&margin=%5Bobject%20Object%5D&name=image.png&originHeight=514&originWidth=1192&originalType=binary&ratio=1&rotation=0&showTitle=false&size=119025&status=done&style=none&taskId=u191bd5b7-c43f-42d6-9ad9-c1e152360d5&title=&width=1192)
微服务的组件增多、架构复杂，使得我们运维变得更加复杂。对于大厂而言，人多维护起来肯定没什么太大问题，可以自建完整的基础设施，但对于小厂而言，研发资源有限，想自建会相当困难~~​~~。
​

不过微服务的基础设施维护困难的问题在 Kubernetes 出现后逐渐出现了转机。在2014年6月Google开源了Kubernetes后，经过这几年的发展，已逐渐成为容器编排领域的事实标准。同时 Kubernetes 已俨然成为云原生时代的超级操作系统，它使得基础设施维护变得异常~~​~~简单。
~~​~~

在传统模式下，我们不仅需要关注应用开发阶段存在的问题，同时还需要关心应用的测试、编译、部署、观测等问题，例如程序是使用systemd、supervisor启动、还是写bash脚本启动？日志是如何记录、如何采集、如何滚动？我们如何对服务进行观测？Metrics 指标如何采集？采集后的指标如何展示？服务如何实现健康检查、存活检查？服务如何滚动更新？如何对流量进行治理，比如实现金丝雀发布、流量镜像？诸如此类的问题。我们业务代码没写几行，全在考虑和权衡基础设施问题。然而~~​~~使用Kubernetes后，可以发现大部分问题都已经被Kubernetes或周边的生态工具解决了，我们仅仅只需要关心上层的应用开发和维护Kubernetes集群即可。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631261444256-a8a6d376-f89f-4618-bee9-318d5b08ab40.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=464&id=UyWn3&margin=%5Bobject%20Object%5D&name=image.png&originHeight=464&originWidth=1173&originalType=binary&ratio=1&rotation=0&showTitle=false&size=80190&status=done&style=none&taskId=u9bac8cb3-04cf-4065-8456-08bfb7ce023&title=&width=1173)
Kubernetes在微服务中的作用就如同建高楼的地基，做了很多基础工作，统一了大量的基础设施标准，以前我们要实现服务的启动、配置、日志采集、探活等功能需要写很多中间件，现在我们只需要写写yaml文件，就可以享受这些基础设施的能力。运维更加简单这个也显而易见，例如在以前出现流量高峰时研发提工单后增加副本数，运维处理工单，人肉扩缩容，现在我们可以根据实际应用的负载能力，合理的配置好副本 CPU、Mem 等资源及 HPA 规则，在流量高峰时由 Kubernetes 自动扩容、流量低谷时自动缩容，省去了大量人工操作。
​

~~​~~同时在框架层面，传统模式下基础设施组件很多都是自研的，基本上没有太多标准可言，框架需要做各种switch case对这种基础设施组件的适配，并且框架经常会为因为基础设施的改变，做一些不兼容的升级。现在只需要适配Kubernetes即可，大大简化微服务的框架难度和开发成本。


​

## 2 微服务的生命周期
刚才我们讲到Kubernetes的优势非常明显，在这里会描述下我们自己研发的微服务框架Ego怎么和Kubernetes结合起来的一些有趣实践。
​

我们将微服务的生命周期分为以下6个阶段：开发、测试、部署、启动、调用、治理。
### ![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631261281413-347cec08-03e5-4195-b19b-cfcd4ed115dc.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=145&id=u7f5f289e&margin=%5Bobject%20Object%5D&name=image.png&originHeight=145&originWidth=725&originalType=binary&ratio=1&rotation=0&showTitle=false&size=15939&status=done&style=none&taskId=ue90bbaf6-b8f8-4743-963a-284387547a6&title=&width=725)
### 2.1 开发阶段
在开发阶段我们最关注三个问题。 如何配置、如何对接，如何~~​~~调试。
​

#### 2.1.1 配置驱动
大家在使用开源组件的时候，其实会发现每个开源组件的配置、调用方式、debug方式、记录日志方式都不一样，导致我们需要不停去查看组件的示例、文档、源码，才能使用好这个组件。我们只想开发一个功能，却需要关心这么多底层实现细节，这对我们而言~~​~~是一个很大的心智负担。


所以我们将配置、调用方式做了统一。可以看到上图我们所有组件的地址都叫addr，然后在下图中我们调用redis、gRPC、MySQL的时候，只需~~要​~~基于组件的配置Key path去 Load 对应的组件配置，通过build方法就可以构造一个组件实例。可以看到调用方式完全相同，就算你不懂这个组件，你只要初始化好了，就可以根据编辑器代码提示，调用这个组件里的API，大大简化我们的开发流程。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631261241446-dcab1c73-859b-44c2-a895-792673154432.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=459&id=u545d0e96&margin=%5Bobject%20Object%5D&name=image.png&originHeight=459&originWidth=867&originalType=binary&ratio=1&rotation=0&showTitle=false&size=113628&status=done&style=none&taskId=u18a824a7-05ba-45ef-9872-8387fb51c4f&title=&width=867)
​

#### 2.1.2 配置补齐
配置补齐这个功能，是源于我们在最开始使用一些组件库的时候，很容易遗漏配置，例如使用`gRPC`的客户端，未设置连接错误、导致我们在阻塞模式下连接不上的时候，没有报正确的错误提示；或者在使用Redis、MySQL没有超时配置，导致线上的调用出现问题，产生雪崩效应。这些都是因为我们对组件的不熟悉，才会遗漏配置。框架要做的是在用户不配置的情况下，默认补齐这些配置，并给出一个最佳实践配置，让业务方的服务更加稳定、高效。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631261307664-35cb9cac-41df-4f45-8511-6efb6e78266c.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=471&id=ube89a900&margin=%5Bobject%20Object%5D&name=image.png&originHeight=471&originWidth=582&originalType=binary&ratio=1&rotation=0&showTitle=false&size=85447&status=done&style=none&taskId=u6b915c16-bdd3-44b4-bc6b-d5ece4c12b1&title=&width=582)
#### 2.1.3 配置工具
我们编写完配置后，需要将配置发布到测试环境，我们将配置中心IDE化，能够非常方便的编写配置，通过鼠标右键，就可以插入资源引用，鼠标悬停可以看到对应的配置信息。通过配置中心，使我们在对比配置版本，发布，回滚，可以更加方便。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631262349303-8d7ba434-7c3e-4b09-8326-186192a1b256.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=528&id=u7efb695b&margin=%5Bobject%20Object%5D&name=image.png&originHeight=528&originWidth=867&originalType=binary&ratio=1&rotation=0&showTitle=false&size=136900&status=done&style=none&taskId=u7236f1f5-b375-4ed1-8989-b0f1c853364&title=&width=867)
#### 2.1.4 对接-Proto管理
我们内部系统全部统一采用`gRPC`协议和`protobuf`编解码。统一的好处在于不需要在做任何协议、编解码转换，这样就可以使我们所有业务采用同一个`protobuf`仓库，基于 CI/CD 工具实现许多自动化功能。
​

我们要求所有服务提供者提前在独立的路径下定义好接口和错误码的protobuf文件，然后提交到GitLab，我们通过GitLab CI的check阶段对变更的protobuf文件做format、lint、breaking 检查。然后在build阶段，会基于 protobuf 文件中的注释自动产生文档，并推送至内部的微服务管理系统接口平台中，还会根据protobuf文件自动构建  Go/PHP/Node/Java 等多种语言的桩代码和错误码，并推送到指定对应的中心化仓库~~​~~。
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631262970245-eaa4f021-b177-44c2-97f4-36ca5f559c02.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=872&id=u4d28f4ff&margin=%5Bobject%20Object%5D&name=image.png&originHeight=872&originWidth=2560&originalType=binary&ratio=1&rotation=0&showTitle=false&size=530378&status=done&style=none&taskId=u524c77c7-787a-4658-b675-6ee68a1e14c&title=&width=2560)推送到仓库后，我们就可以通过各语言的包管理工具拉取客户端、服务端的gRPC和错误码的依赖，不需要口头约定对接数据的定义，也不需要通过 IM 工具传递对接数据的定义文件，极大的简化了对接成本~~​~~。
#### 
#### 2.1.5 对接-错误码管理
有了以上比较好的protobuf生成流程后，我们可以进一步简化业务错误状态码的对接工作。而我们采用了以下方式：~~​~~

- Generate：
    - 编写protobuf error的插件，生成我们想要的error代码
    - 根据go官方要求，实现errors的interface，他的好处在于可以区分是我们自定义的error类型，方便断言。

![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631805432392-23326c11-5be3-4295-9e79-502195d2829e.png#clientId=ue5f57623-9a54-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=208&id=u2c4fe109&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1087&originWidth=2560&originalType=binary&ratio=1&rotation=0&showTitle=false&size=435006&status=done&style=none&taskId=u766edfdb-1627-4483-9a5c-155971186ec&title=&width=491)

- 根据注解的code信，在错误码中生成对应的grpc status code，业务方使用的时候少写一行代码

![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631806163171-cc166add-6807-47fc-8d78-41715905193d.png#clientId=ue5f57623-9a54-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=326&id=ub02b6816&margin=%5Bobject%20Object%5D&name=image.png&originHeight=326&originWidth=257&originalType=binary&ratio=1&rotation=0&showTitle=false&size=25655&status=done&style=none&taskId=u4101a16a-1413-415c-bc41-8ae3de9d3dd&title=&width=257)![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631806141251-4a9eb1a2-dd47-469f-b3fa-42e57559a93c.png#clientId=ue5f57623-9a54-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=153&id=JvMMo&margin=%5Bobject%20Object%5D&name=image.png&originHeight=160&originWidth=476&originalType=binary&ratio=1&rotation=0&showTitle=false&size=31414&status=done&style=none&taskId=u0667acad-31d0-4a62-a3f1-d1b0c8c44ed&title=&width=455)

- 确保错误码唯一，后续在API层响应用户数据确保唯一错误码，例如: 下单失败(xxx)
- errors里设置with message，with metadata，携带更多的错误信息
- Check：
    - gRPC的error可以理解为远程error，他是在另一个服务返回的，所以每次error在客户端是反序列化，new出来的。是无法通过errors.Is判断其根因。

![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631707435739-7f8c7172-ccb0-4462-8021-49e97478d88e.png#clientId=uad1fbcbd-cca7-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=256&id=u63b9fb45&margin=%5Bobject%20Object%5D&name=image.png&originHeight=256&originWidth=607&originalType=binary&ratio=1&rotation=0&showTitle=false&size=26418&status=done&style=none&taskId=u3b09419b-7694-4e4b-be9d-eaa1606db05&title=&width=607)

- 我们通过工具将gRPC的错误码注册到一起，然后客户端通过FromError方法，从注册的错误码中，根据Reason的唯一性，取出对应的错误码，这个时候我们可以使用errors.Is来判断根因。

![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631267923383-e4356248-1af7-4c97-bd00-ef80ed227871.png#clientId=u9a05a0ae-3d87-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=424&id=uf2397b77&margin=%5Bobject%20Object%5D&name=image.png&originHeight=424&originWidth=1112&originalType=binary&ratio=1&rotation=0&showTitle=false&size=86004&status=done&style=none&taskId=u1c2c3534-f7cf-46a2-be79-684597f113f&title=&width=1112)
![image.png](https://cdn.nlark.com/yuque/0/2021/png/497518/1631806238170-a5adb0a8-2590-48c6-b712-4001c15ae0e4.png#clientId=ue5f57623-9a54-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=1440&id=u712da689&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1440&originWidth=2214&originalType=binary&ratio=1&rotation=0&showTitle=false&size=502482&status=done&style=none&taskId=u5456a670-d2cb-414c-a42e-d2f94e0b746&title=&width=2214)

- 最后做到errors.Is的判断: errors.Is(eerrors.FromError(err), UserErrNotFound())
#### 
#### 2.1.6 对接-调试
对接中调试的第一步是阅读文档，我们之前通过protobuf的ci工具里的lint，可以强制让我们写好注释，这可以帮助我们生成非常详细的文档。

基于 gRPC Reflection 方法，服务端获得了暴露自身已注册的元数据能力，第三方可以通过 Reflection 接口获取服务端的 Service、Message 定义等数据。结合 Kubernetes API，用户选择集群、应用、Pod 后，可直接在线进行gRPC接口测试。同时我们可以对测试用例进行存档~~​~~，方便其他人来调试该接口。
