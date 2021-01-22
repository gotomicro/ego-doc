# 阿里云日志查询语句
### 1 gRPC服务端日志
#### 1.1 获取gRPC请求应用分布
```sql
* and comp:"server.egrpc" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 1.2 获取gRPC请求的应用状态码分布
```sql
* and comp:"server.egrpc" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 1.3 获取gRPC请求的应用崩溃分布
```sql
* and comp:"server.egrpc" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 1.4 获取gRPC请求的应用错误分布
```sql
* and comp:"server.egrpc" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 1.5 获取gRPC请求的应用慢日志分布
```sql
* and comp:"server.egrpc" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 1.6 获取某个应用gRPC的对端地址的错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"server.egrpc" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, peerIp GROUP by peerIp order by count desc limit 10
```
#### 1.7 获取某个应用gRPC的对端应用名的错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"server.egrpc" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, peerName GROUP by peerName order by count desc limit 10
```

### 2 HTTP服务端日志
#### 2.1 获取HTTP请求应用分布
```sql
* and comp:"server.ehttp" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 2.2 获取HTTP请求的应用状态码分布
```sql
* and comp:"server.ehttp" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 2.3 获取HTTP请求的应用崩溃分布
```sql
* and comp:"server.ehttp" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 2.4 获取HTTP请求的应用错误分布
```sql
* and comp:"server.ehttp" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 2.5 获取HTTP请求的应用慢日志分布
```sql
* and comp:"server.ehttp" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 2.6 获取某个应用HTTP的对端地址的错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"server.ehttp" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, peerIp GROUP by peerIp order by count desc limit 10
```
#### 2.7 获取某个应用HTTP的对端应用名的错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"server.ehttp" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, peerName GROUP by peerName order by count desc limit 10
```


### 3 gRPC客户端日志
#### 3.1 获取gRPC请求应用分布
```sql
* and comp:"client.egrpc" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 3.2 获取gRPC请求的应用状态码分布
```sql
* and comp:"client.egrpc" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 3.3 获取gRPC请求的应用崩溃分布
```sql
* and comp:"client.egrpc" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 3.4 获取gRPC请求的应用错误分布
```sql
* and comp:"client.egrpc" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 3.5 获取gRPC请求的应用慢日志分布
```sql
* and comp:"client.egrpc" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 3.6 获取某个应用gRPC的配置项错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"client.egrpc" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, compName GROUP by compName order by count desc limit 10
```


### 4 HTTP客户端日志
#### 4.1 获取HTTP请求应用分布
```sql
* and comp:"client.ehttp" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 4.2 获取HTTP请求的应用状态码分布
```sql
* and comp:"client.ehttp" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 4.3 获取HTTP请求的应用崩溃分布
```sql
* and comp:"client.ehttp" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 4.4 获取HTTP请求的应用错误分布
```sql
* and comp:"client.ehttp" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```

#### 4.5 获取HTTP请求的应用慢日志分布
```sql
* and comp:"client.ehttp" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```

#### 4.6 获取某个应用HTTP 的配置项错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"client.ehttp" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, compName GROUP by compName order by count desc limit 10
```

### 5 Gorm客户端日志
#### 5.1 获取Gorm请求应用分布
```sql
* and comp:"client.egorm" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 5.2 获取Gorm请求的应用状态码分布
```sql
* and comp:"client.egorm" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 5.3 获取Gorm请求的应用崩溃分布
```sql
* and comp:"client.egorm" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 5.4 获取Gorm请求的应用错误分布
```sql
* and comp:"client.egorm" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 5.5 获取Gomr请求的应用慢日志分布
```sql
* and comp:"client.egorm" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 5.6 获取某个应用Gorm的配置项错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"client.egorm" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, compName GROUP by compName order by count desc limit 10
```


### 6 Redis客户端日志
#### 6.1 获取Redis请求应用分布
```sql
* and comp:"client.eredis" and msg:"access"  | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 6.2 获取Gorm请求的应用状态码分布
```sql
* and comp:"client.eredis" and msg:"access" | SELECT  COUNT(*) as count, code GROUP by code order by count desc limit 10
```
#### 6.3 获取Redis请求的应用崩溃分布
```sql
* and comp:"client.eredis" and msg:"access" and lv:"panic" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 6.4 获取Redis请求的应用错误分布
```sql
* and comp:"client.eredis" and msg:"access" and lv:"error" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```

#### 6.5 获取Redis请求的应用慢日志分布
```sql
* and comp:"client.eredis" and msg:"access" and event:"slow" | SELECT  COUNT(*) as count, app GROUP by app order by count desc limit 10
```
#### 6.6 获取某个应用Redis的配置项错误分布
其他崩溃日志、慢日志与该语句类似
```sql
* and comp:"client.eredis" and msg:"access" and lv:"error" and app:"appname" | SELECT  COUNT(*) as count, compName GROUP by compName order by count desc limit 10
```

