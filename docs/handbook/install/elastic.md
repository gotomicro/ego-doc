## 安装elastic
### 下载
```bash
docker pull elasticsearch:7.13.2
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.13.2
```
### 测试
```bash
curl http://127.0.0.1:9200/     
{
  "name" : "c87ed7c2e927",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "YVwG5FiWRrSoF4uGUqdyBA",
  "version" : {
    "number" : "7.13.2",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "4d960a0733be83dd2543ca018aa4ddc42e956800",
    "build_date" : "2021-06-10T21:01:55.251515791Z",
    "build_snapshot" : false,
    "lucene_version" : "8.8.2",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

## 安装kibana
### 下载
```bash
docker pull kibana:7.13.2
```

### 启动kibana
```bash
docker run --name kibana -e ELASTICSEARCH_HOSTS=http://172.16.21.254:9200 -p 5601:5601 -d kibana:7.13.2
```

### 测试kibana
```bash
docker logs -f kibana
curl http://127.0.0.1:5601
```

## 常用elastic
```bash
curl -k -u admin:admin -XGET http://127.0.0.1:9200/_cluster/health\?pretty
{
  "cluster_name" : "docker-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 6,
  "active_shards" : 6,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```

```bash
curl -k -u admin:admin 'http://127.0.0.1:9200/_cat/indices?v'
health status index                           uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .kibana_7.13.2_001              Wa5NDmYRSaWIh08S0yrTPA   1   0         21           13      2.1mb          2.1mb
green  open   .apm-custom-link                hcERyl5TTIWPil1U4XgHlQ   1   0          0            0       208b           208b
green  open   .kibana-event-log-7.13.2-000001 pzUQqkfwSuSXQl-MRbIZ-w   1   0          1            0      5.6kb          5.6kb
green  open   .apm-agent-configuration        hOW933KgTjKcitUG4YLPeg   1   0          0            0       208b           208b
green  open   .kibana_task_manager_7.13.2_001 brBQqHiDSAeDMQOrTRLduQ   1   0         10          811    156.3kb        156.3kb
```

```bash
curl -k -u admin:admin -XPUT http://127.0.0.1:9200/ego_logger 
curl -k -u admin:admin -XPOST http://127.0.0.1:9200/ego_logger/product/ -H "Content-Type: application/json" -d '{"name":"ego", "author": "ego logger", "c_version": "2.7.3"}' 
```