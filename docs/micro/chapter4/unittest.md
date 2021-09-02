## 单元测试

### 概念

单元测试是指对软件中最小可测试单元在于程序其他部分相隔离的情况下进行检查和校验。单元测试是白盒测试，通常由开发来完成，测试的对象可以是函数，也可以是类，而测试的目标是来检查程序是否按照预期的逻辑执行。好的单元测试会遵守 AIR
原则，以便让测试用例更加地复合规范。

### AIR原则

AIR，即 Automatic、Indepenndent、Repleatable 的简写：

* Automatic：自动化，单元测试应该是自动执行，自动校验，自动给出结果，若需要人工检查（如将结果输出到控制台）的单元测试，不是好的单元测试；
* Indepenndent：独立性，单元测试应该是可以独立运行的，测试用例之间无依赖和执行次序，用例内部对外部资源也无依赖；
* Repleatable：可重复，单元测试应该是可以重复执行的，每次的结果都是稳定可靠的；

### 如何简单高效做单元测试

目前单元测试大部分的玩法，都是在做解除依赖，例如以下的一些方式

* 面向接口编程
* 依赖注入、控制反转
* 使用Mock

不可否认，以上的方法确实可以使代码变得更加优雅，更加方便测试。但是实现了以上的代码，会让我们的代码变得更加复杂、增加更多的开发工作量，下班更晚。

如果我们不方便解除依赖，我们是否可以让基础设施将所有依赖构建起来。基础设施能做的事情，就不要让研发用代码去实现。

以下举我们一个实际场景的`MySQL`单元测试例子。我们可以通过`gitlab.yaml`，构建一个`mysql`。然后通过`Ego`的应用执行`job`。

* 创建数据库的表`./app --job=install`
* 初始化数据库表中的数据 `./app --job=intialize`
* 执行`go test ./...`

通过`CI`中构建`MySQL`，在定义公司内部的`install`和`initialize`标准，我们能够很方便的将依赖的数据源构建起来，方便做测试。

CI怎么玩，后面在介绍

### HTTP 接口级别单元测试

[HTTP接口测试Example](https://github.com/gotomicro/ego-doc/blob/main/examples/unittest/http/main_test.go)

### gRPC 接口级别单元测试

[gRPC接口测试Example](https://github.com/gotomicro/ego-doc/blob/main/examples/unittest/grpc/main_test.go)

### 引用
* [Golang单元测试实践指南](https://mp.weixin.qq.com/s?src=11&timestamp=1630576331&ver=3290&signature=BcOBn5vIjb9uNyOe0162iqGckbT6CfTNSI2KR5EKTKRYmFPyajDLlGMBIdaYo91FcVraCTTH71BcgHAjL2nVF76p4UJYXvto5uWAqnTWLwFW7d6F15C-HkpFjGH90myW&new=1)
* [如何做好单元测试](https://mp.weixin.qq.com/s?src=11&timestamp=1630576201&ver=3290&signature=PLyPvRKKk25r22J46kIgp08IyO3R8iQW*Y6KKdTM1zsYTheyMAoTZmKtiNijpVGVujRUGAshDu64ldzGjmMoifSnADzgH-yV5rsQM4fPOIYM4Jj2cnB5F7rB97ONwegk&new=1)
* [分享一个 UT failed 引出的思考](https://mp.weixin.qq.com/s/iH4JdOCybwfaANT97zW7cg)
