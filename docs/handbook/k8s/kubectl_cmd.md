# kubectl常用指令
## 查看某个pod情况
```bash
kubectl describe pod {POD}
```

## 进入pod
```bash
kubectl exec -it {POD} sh
```

## 创建secret
```bash
kubectl create secret docker-registry {key} \
--docker-server={server} --docker-username={username} --docker-password={password}
```

## 创建role
```bash
 kubectl create role pods-reader --verb=get,list,watch --resource=pods,endpoints --dry-run -o yaml
 kubectl create rolebinding app-api-default-pods-reader --role=pods-reader --namespace=default --dry-run -o yaml > rolebinding.yaml
```
kubectl  get rolebinding -o yaml app-api-default-pods-reader
kubectl describe rolebinding app-api-default-pods-reader

## 查看某个角色
```bash
kubectl get role
```

## 查看角色情况
```bash
kubectl describe role pods-reader
```

## 执行yaml
```bash
kubectl apply -f role-reader.yaml
```


kubectl  get rolebinding -o yaml default-sa-resource-reader-binding
kubectl  get role -o yaml resource-reader


# 集群级别的  角色定义和授权:
kubectl get ClusterRole  -o yaml cluster-admin
kubectl get clusterRolebinding -o yaml admin


kubectl logs -f user-svc-dp-75985d755b-gntkz


kubectl get ep -o yaml