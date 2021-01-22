# 监控查询语句
## 1 全部请求量
* 全部请求量 sum(irate(ego_server_handle_total{}[1m]))
* 某个应用请求量 sum(irate(ego_server_handle_total{app="你的应用名称"}[1m]))