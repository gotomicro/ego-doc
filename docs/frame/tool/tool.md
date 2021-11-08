# 全家桶工具
## example
EGO服务端生成PB，生成单元测试例子：

https://github.com/gotomicro/go-engineering/tree/main/chapter_ego_unittest

## 下载工具
使用bash脚本下载工具
```bash
bash <(curl -L https://raw.githubusercontent.com/gotomicro/egoctl/main/getlatest.sh)
```
通过以上脚本，可以下载protoc工具全家桶，以及ego的protoc插件和egoctl
* /usr/local/bin/protoc
* /usr/local/bin/protoc-gen-go
* /usr/local/bin/protoc-gen-go-grpc
* /usr/local/bin/protoc-gen-openapiv2
* /usr/local/bin/protoc-gen-go-errors
* /usr/local/bin/protoc-gen-go-http
* /usr/local/bin/egoctl

## 软连接和生成测试代码 
```makefile
PROTO:=protos
PROJECT_NAME=helloworld

# 挂载Proto
link:export PROTO_DIR=../chapter_proto
link:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>make $@<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@ if [[ ! -L $(APP_PATH)/$(PROTO) ]]; then ln -s $(PROTO_DIR) $(APP_PATH)/$(PROTO); echo "link created"; else echo "link exists"; fi;

# 生成pb
gen-proto:
	@protoc -I ./$(PROTO) --go_out=paths=source_relative:./$(PROTO) ./$(PROTO)/$(PROJECT_NAME)/*.proto
	@protoc -I ./$(PROTO) --go-grpc_out=paths=source_relative:./$(PROTO) ./$(PROTO)/$(PROJECT_NAME)/*.proto
	@protoc -I ./$(PROTO) --go-errors_out=paths=source_relative:./$(PROTO) ./$(PROTO)/$(PROJECT_NAME)/*.proto

# 生成单元测试
gen-test:
	@protoc -I ./$(PROTO)  --go-test_out=out=./server/router,paths=source_relative:. ./$(PROTO)/$(PROJECT_NAME)/*.proto
```

## 使用原理 
### 生成PB文件
```bash
protoc -I {你的proto路径} --go_out=paths=source_relative:{输出路径} {你的proto来源文件}
```
### 生成gRPC文件
```bash
protoc -I {你的proto路径} --go-grpc_out=paths=source_relative:{输出路径} {你的proto来源文件}
```
### 生成errors文件
```bash
protoc -I {你的proto路径} --go-grpc_out=paths=source_relative:{输出路径} {你的proto来源文件}
```
### 生成单元测试文件
```bash
protoc -I {你的proto路径}  --go-test_out=out={输出到指定目录},paths=source_relative:{输出路径} {你的proto来源文件}
```