# GRPC-Gateway
* [ EGO代码实现 ] 代码层面利用反射实现gRPC转成HTTP：https://github.com/gotomicro/ego/tree/master/examples/http/grpc-proxy
* [ EGO网关实现 ] 通过gRPC反射接口和gRPC resolver动态发现，实现gateway: https://github.com/gotomicro/ego-gateway
* 通过protobuf编译生成HTTP的gw.go的文件，编译到gateway中 https://github.com/grpc-ecosystem/grpc-gateway
* 通过提供API接口，接收proto等元数据信息，类似于apisix方式，将grpc转成http：https://github.com/apache/apisix/blob/master/docs/en/latest/plugins/grpc-transcode.md


