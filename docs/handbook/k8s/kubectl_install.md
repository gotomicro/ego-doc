# 安装kubectl
## 安装指令
```bash
brew install kubectl
brew install kubectx
brew install kubens
```

### 安装kube
如果brew安装慢，是因为网络问题，我们可以手动下载。
https://hub.fastgit.org/ahmetb/kubectx/releases

### 安装配置
创建`.kube`目录
```bash
mkdir ~/.kube
```

将k8s配置加入到`~/.kube/config`
```yaml
apiVersion: v1
clusters:
  - cluster:
        server: [K8S API地址]
        certificate-authority-data: [认证数据]
    name: [集群名字]
contexts:
  - context:
        cluster: [集群名字]
        user: [集群用户名]
        namespace: default
    name: [集群名字]
kind: Config
preferences: {}
users:
  - name: [集群用户名]
    user:
        client-certificate-data: [认证数据]
        client-key-data: [认证key]
```

### 测试指令
```bash
# 查看集群列表
➜  ~ kubectx
[集群名字]

# 选择集群名称
➜  ~ kubectx [集群名字]
✔ Switched to context [集群名字].

# 查看命名空间列表
➜  ~ kubens           
default
kube-node-lease
kube-public
kube-system

# 选择命名空间
➜  ~ kubens default
✔ Active namespace is "default"

# 查看Pod信息
➜  ~ kubectl get pod                              
NAME                    READY   STATUS    RESTARTS   AGE
test-6d565dbb89-9456x   1/1     Running   0          16h
```