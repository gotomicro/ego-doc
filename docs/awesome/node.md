## Node监控
## TCP状态
```shell
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'    
```

## Sockstat
```shell
cat /proc/net/sockstat

sockets: used 294
TCP: inuse 35 orphan 0 tw 0 alloc 45 mem 1
UDP: inuse 13 mem 2
UDPLITE: inuse 0
RAW: inuse 4
FRAG: inuse 0 memory 0
```
### 含义
sockets: used：已使用的所有协议套接字总量
TCP: inuse：正在使用（正在侦听）的TCP套接字数量。其值≤ netstat –lnt | grep ^tcp | wc –l
TCP: orphan：无主（不属于任何进程）的TCP连接数（无用、待销毁的TCP socket数）
TCP: tw：等待关闭的TCP连接数。其值等于netstat –ant | grep TIME_WAIT | wc –l
TCP：alloc(allocated)：已分配（已建立、已申请到sk_buff）的TCP套接字数量。其值等于netstat –ant | grep ^tcp | wc –l
TCP：mem：套接字缓冲区使用量（单位不详。用scp实测，速度在4803.9kB/s时：其值=11，netstat –ant 中相应的22端口的Recv-Q＝0，Send-Q≈400） UDP：inuse：正在使用的UDP套接字数量 RAW： FRAG：使用的IP段数量

### Prometheus采集的sockstat metric
node_sockstat_TCP_tw
node_sockstat_TCP_mem_bytes

## 


