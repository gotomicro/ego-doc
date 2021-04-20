# Release Note
## v0.5.4
* 正式环境下，panic时候，文件和命令行双输出。并且将field信息，高亮输出到终端。
* 调试模式下，http，gRPC客户端发送请求，会输出行号，该行号可以直接用goland打开。
* resolver阶段，无锁化判断，attributes信息是否变更
* 修改readme文档，描述客户端和服务端gRPC链路玩法
* 优化gRPC客户端设置metadata信息的性能
* 根据客户端头文件的控制，返回gRPC服务端的CPU利用率
* 支持gRPC流式客户端设置metadata信息
* panic增加配置名
