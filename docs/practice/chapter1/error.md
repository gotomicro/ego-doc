## 唯一状态码

错误码是识别业务系统的核心手段。
* 错误码的生成
* 如何记录错误码
* 如何排查错误码

### 1 错误码的生成
我们使用`proto`约束我们的错误码。并使用`ego`中的`protoc-gen-go-errors`插件将我们在`proto`中定义的@code，@i18n转为对应的代码。

```proto
syntax = "proto3";
package biz.v1;
// 下行注解是必要的，有这个注解 protoc-gen-go-errors 才尝试解析当前 protobuf 文件中的 enum，并基于 enum 生成错误桩代码
// @plugins=protoc-gen-go-errors
// language-specified package name
option go_package = "biz/v1;bizv1";
option java_multiple_files = true;
option java_outer_classname = "BizProto";
option java_package = "com.ego.biz.v1";
// enum Err 定义了错误的不同枚举值，protoc-gen-go-errors 会基于 enum Err 枚举值生成错误桩代码
// @code 为错误关联的gRPC Code (遵循 https://grpc.github.io/grpc/core/md_doc_statuscodes.html 定义，需要全大写)，
//       包含 OK、UNKNOWN、INVALID_ARGUMENT、PERMISSION_DENIED等
// @i18n.cn 国际化中文文案
// @i18n.en 国际化英文文案
enum Err {
  // 请求正常，实际上不算是一个错误
  // @code=OK
  // @i18n.cn="请求成功"
  // @i18n.en="OK"
  ERR_OK = 0;
  // 未知错误，比如业务panic了
  // @code=UNKNOWN             # 定义了这个错误关联的gRPC Code为：UNKNOWN
  // @i18n.cn="服务内部未知错误" # 定义了一个中文错误文案
  // @i18n.en="unknown error"  # 定义了一个英文错误文案
  ERR_UNKNOWN = 1;
  // 找不到指定用户
  // @code=NOT_FOUND
  // @i18n.cn="找不到指定用户"
  // @i18n.en="user not found"
  ERR_USER_NOT_FOUND = 2;
}
```
我们会生成这样的错误码
```go
// ErrUnknown  未知错误，比如业务panic了
// @code=UNKNOWN             # 定义了这个错误关联的gRPC Code为：UNKNOWN
// @i18n.cn="服务内部未知错误"        # 定义了一个中文错误文案
// @i18n.en="unknown error"  # 定义了一个英文错误文案
func ErrUnknown() eerrors.Error {
	return errUnknown
}
```

如果我们需要国际化的错误信息
```go
ReasonI18n(ErrUnknown(), "cn")
```

### 2 如何记录错误码
我们遇到错误，需要做两个处理，第一将错误码通过接口反馈给用户，第二将错误码记录到日志里。
第一个比较好处理。在这里不在阐述。
将错误码记录到日志里，我们有四点要注意
* 错误码不要记录到业务日志里，因为业务日志是没有规范的，很难去检索和排查。因此我们最好的方式，是将错误记录到框架日志里
* 错误码记录到日志里，同时还需要记录对应的详细Error信息
* Error不能够吞掉，需要Wrap整个链路
* 错误码的唯一性

#### 2.1 如何记录错误码到框架日志
框架日志追加字段，是`EGO`框架的非常重要的特性。我们可以在环境变量里设置你要需要追加的字段，比如`export EGO_LOG_EXTRA_KEYS=X-Ego-Code,X-Ego-Uid`
`EGO`会提取Context中这些字段信息，追加到对应的框架日志里。这样我们就可以根据框架里的`method`,`ucode`,`msg`,`comp`等多个框架字段和错误码配合使用，分析处更多错误信息。
在`HTTP`服务中使用方法
```go
parentContext := transport.WithValue(ctx.Request.Context(), "X-Ego-Code", proto.ERR_USER_NOT_FOUND)
ctx.Request = ctx.Request.WithContext(parentContext)
```

`gRPC`服务会比较特殊。在用户代码里无法向`ctx`中写入数据，因此框架的拦截器无法获取数据。因此我们需要使用`CtxStoreSet`方法
```go
egrpc.CtxStoreSet(ctx, "X-Ego-Code", proto.ERR_USER_NOT_FOUND)
```

#### 2.2 如何记录Error到框架日志
我们遇到error错误，会习惯性的看到一个error就立刻记录一个错误日志，这种方式并不是最有效。假设一个场景。一个controller调用一个mysql出错。mysql模块记录了这个error日志，controller也记录了这个error日志，这样我们就会有两条日志，杂音太多。
因此我们需要`Wrap`整个链路。这样我们就可以在最顶层记录一条日志。

然后我们可以在`HTTP`的`gin`服务里使用`ctx.SetErr(err)方法`，`EGO`会自动将该错误记录到框架日志的`err`字段里。
在`gRPC`的服务里只要你响应的错误里不为`nil`，`EGO`也会自动将该错误记录到框架日志的`err`字段里。

#### 2.3 错误码的唯一性
错误码的唯一性是非常必要的。如果大家看过一些工厂的设备，由很长的错误码，每个错误码都代表一种解决方案。这样我们就可以根据错误码，快速定位问题。

设想，如果我们一个错误码代表了n种原因，我们就需要再进一步追查别的信息，定位真正的原因。这样就会增加我们的排查成本。

如何保证他的唯一性？因为我们的错误码都是统一管理，我们可以利用ast分析多个仓库，确保没有错误码被使用两次，同时也可以识别出错误码是否被使用。

### 3 如何排查错误码
我们可以直接在日志系统里使用
`x-shimo-code='ERR_USER_NOT_FOUND'`，就可以看到


