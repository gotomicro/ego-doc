## 镜像
登录 https://registry.hub.docker.com/

## 常用镜像
* [编译Go](https://registry.hub.docker.com/_/golang)
```go
docker pull golang:1.16.4-buster
```
* [linux运行镜像](https://registry.hub.docker.com/_/alpine)
```bash
docker pull alpine:3.13.5
```

## 自己构建镜像
通常我们会要调试一些代码，所以需要自己构建镜像
* [centos版本](https://registry.hub.docker.com/_/centos)
```bash
docker pull centos:7.9.2009
```

例如我们想装个GO环境调试一些问题

可以在`golang.org`下载一个linux的Go环境，然后编写一个`Dockerfile`

```dockerfile
FROM centos:7.9.2009
RUN yum -y install wget git vim net-tools gcc automake autoconf libtool make gcc-c++
ADD go /root/go
ENV GOPATH=/code/go
ENV GOROOT=/root/go
ENV PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
CMD ["/sbin/my_init"]
```

```bash
docker build -t [你的GO调试镜像] ./
docker run -itd  --name gotestrun [你的GO调试镜像]  /usr/sbin/init
docker exec -it gotestrun /bin/bash
# 进入到你的容器
go version
```
