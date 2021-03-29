# 2 编译

Go常用部署方式有三种 `nohup` 、 `supervisor` 、 `systemd` 。

## 2.1 Nohup
nohup是在linux系统上运行程序的一种方式。这种方式比较简单，通过一个命令行，就可以快速启动程序
```bash
nohup ./micro &
```

- nohup 加在一个命令的最前面，表示不挂断的运行命令
- & 加在一个命令的最后面，表示这个命令放在后台执行
### 2.1.1 常用指令

- 查看进程 jobs
```bash
➜  ~ jobs
[1]  + running    nohup ./micro
```
“+”代表最近的一个任务（当前任务），“-”代表之前的任务。
只有在当前命令行中使用 nohup和& 时，jobs命令才能将它显示出来。如果将他们写到 .sh 脚本中，然后执行脚本，是显示不出来的

- 关闭进程
    - 通过jobs命令查看jobnum，然后执行   kill %jobnum
    - 通过ps命令查看进程号PID，然后执行  kill %PID
- 进程切换到前台
    - fg指令  将程序从后台中调至前台继续运行
```bash
➜  ~ fg %1
[1]  + 9037 running    nohup ./micro
```

- ctrl+z 指令 将程序从前台调至后台继续运行
```bash
➜  ~ fg %1
[1]  + 9037 running    nohup ./micro
^Z
[1]  + 9037 suspended  nohup ./micro
```
`nohup` 只能实现程序在后台运行，作用比较弱。不能对程序的panic做监控报警和自动拉起。一般在生产环境上，我们并不会使用 `nohup` ，我们会使用更


## 2.2 Supervisor
### 2.2.1 安装supervisor
```bash
yum install python-setuptools

easy_install supervisor
```
### 2.2.2 测试是否安装成功
```bash
测试是否安装成功
echo_supervisord_conf
```
### 2.2.3 配置supervisor配置
```bash
[program:micro]
directory=/home/micro/
command=/home/micro/micro
process_name=%(program_name)s
user=www
numprocs=1
autostart=true
startsecs=3
startretries=3
autorestart=true
exitcodes=0,2
stopsignal=TERM
stopwaitsecs=10
serverurl=AUTO
stdout_logfile=/home/micro/supervisorlogs/%(program_name)s_stdout.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10
stdout_capture_maxbytes=0
stdout_events_enabled=true
stderr_logfile=/home/micro/supervisorlogs/%(program_name)s_stderr.log
stderr_logfile_maxbytes=50MB
stderr_logfile_backups=10
stderr_capture_maxbytes=0
stderr_events_enabled=false
```
### 2.2.4 执行supervisor指令
```bash
supervisorctl update
supervisorctl start micro
```

