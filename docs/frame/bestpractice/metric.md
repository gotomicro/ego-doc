# 监控查询语句
## 1 系统监控
## 1.1 CPU负载
* 排行榜前10应用 topk(10,sum(rate(process_cpu_seconds_total{}[1m])) by (app))
* 某个应用CPU sum(rate(process_cpu_seconds_total{app="你的应用名称"}[1m]) by (app))
* topk(3, max(max_over_time(irate(process_cpu_seconds_total{}[1m])[1d:])*100)by(job))

## 1 全部请求量
* 全部请求量 sum(irate(ego_server_handle_total{}[1m]))
* 某个应用请求量 sum(irate(ego_server_handle_total{app="你的应用名称"}[1m]))
* 排行榜前10应用 topk(10,sum (rate (ego_server_handle_total{}[1m])) by (app))
* 聚合应用 sum(irate(ego_server_handle_total{}[1m])) by (app)
