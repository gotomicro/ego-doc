## Docker镜像仓库指令
### 登录Docker Registry
```bash
$ sudo docker login --username=[你的名称] [你的Docker Registry地址]
```

### 从Registry中拉取镜像
```bash
$ sudo docker pull [你的Docker Registry地址]/[命名空间]/[镜像名称]:[镜像版本号]
```

### 将镜像推送到Registry
```bash
$ sudo docker login --username=[你的名称] [你的Docker Registry地址]
$ sudo docker tag [ImageId] [你的Docker Registry地址]/[命名空间]/[镜像名称]:[镜像版本号]
$ sudo docker push [你的Docker Registry地址]/[命名空间]/[镜像名称]:[镜像版本号]
```
请根据实际镜像信息替换示例中的[ImageId]和[镜像版本号]参数。


### 示例
使用`docker tag`命令重命名镜像，并将它通过专有网络地址推送至Registry。
```bash
$ sudo docker images
REPOSITORY                                  TAG                 IMAGE ID            CREATED            VIRTUAL SIZE
golang                                      1.16.4-alpine3.13   722a834ff95b        13 days ago        301MB
registry.xxx.com/default/golang:1.16.4      1.16.4-alpine3.13   722a834ff95b        13 days ago        301MB

# 将已有的镜像，打个tag，准备推送远程
$ sudo docker tag 722a834ff95b registry.xxx.com/default/golang:1.16.4
# 使用 "docker push" 命令将该镜像推送至远程。
$ sudo docker push registry.xxx.com/default/golang:1.16.4
```