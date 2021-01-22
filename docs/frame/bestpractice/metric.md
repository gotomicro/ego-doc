# 监控查询语句
## 1 全部请求量
* 全部请求量 sum(irate(ego_server_handle_total{}[1m]))
* 某个应用请求量 sum(irate(ego_server_handle_total{app="你的应用名称"}[1m]))
* 排行榜前10应用 topk(10,sum (rate (ego_server_handle_total{}[1m])) by (app))
* 聚合应用 sum(irate(ego_server_handle_total{}[1m])) by (app)
## 2 CPU负载
* 排行榜前10应用  topk(10,sum(rate(process_cpu_seconds_total{}[1m])) by (app))