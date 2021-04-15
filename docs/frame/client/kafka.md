# Kafka
## 1 Example
[Producer地址](https://github.com/gotomicro/ego-component/tree/master/ekafka/examples)
[Consumer地址](https://github.com/gotomicro/ego-component/blob/master/ekafka/examples/consumerserver)

ego版本：``ego@v0.5.3``

## 2 Producer
### 2.1 Producer配置
```go
type config struct {
    // Brokers brokers地址
    Brokers []string `json:"brokers" toml:"brokers"`
    // Debug 是否开启debug模式
    Debug bool `json:"debug" toml:"debug"`
    // Client 用于创建topic等
    Client clientConfig `json:"client" toml:"client"`
    // Producers 多个消费者，用于生产消息
    Producers map[string]producerConfig `json:"producers" toml:"producers"`
    // Consumers 多个生产者，用于消费消息
    Consumers    map[string]consumerConfig `json:"consumers" toml:"consumers"`
}

type producerConfig struct {
    // Topic 指定生产的消息推送到哪个topic
    Topic string `json:"topic" toml:"topic"`
    // Balancer 指定使用哪种Balancer，可选：hash\roundRobin
    Balancer string `json:"balancer" toml:"balancer"`
    // MaxAttempts 最大重试次数，默认10次
    MaxAttempts int `json:"maxAttempts" toml:"maxAttempts"`
    // BatchSize 批量发送的消息数量，默认100条
    BatchSize int `json:"batchSize" toml:"batchSize"`
    // BatchBytes 批量发送的消息大小，默认1MB
    BatchBytes int64 `json:"batchBytes" toml:"batchBytes"`
    // BatchTimeout 批量发送消息的周期，默认1s
    BatchTimeout time.Duration `json:"batchTimeout" toml:"batchTimeout"`
    // ReadTimeout 读超时
    ReadTimeout time.Duration `json:"readTimeout" toml:"readTimeout"`
    // WriteTimeout 写超时
    WriteTimeout time.Duration `json:"writeTimeout" toml:"writeTimeout"`
    // RequiredAcks ACK配置
    // RequireNone (0) fire-and-forget，producer不等待来自broker同步完成的确认后，就可以发送下一批消息
    // RequireOne  (1) producer在leader已成功收到的数据并得到确认后，才发送下一批消息
    // RequireAll  (-1) producer在所有follower副本确认接收到数据后，才发送下一批消息
    RequiredAcks kafka.RequiredAcks `json:"requiredAcks" toml:"requiredAcks"`
    // Async 设置成true时会导致WriteMessages非阻塞，会导致调用WriteMessages方法获取不到error
    Async bool `json:"async" toml:"async"`
}

```

### 2.2 优雅的Debug
通过开启``debug``配置和命令行的``export EGO_DEBUG=true``，我们就可以在测试环境里看到请求里的配置名、地址、耗时、请求数据、响应数据
![img.png](../../images/frame/client/kafka/kafka1.png)


### 2.3 用户配置
```toml
[kafka]
    debug=true
    brokers=["localhost:9091","localhost:9092","localhost:9093"]
[kafka.client]
    timeout="3s"
[kafka.producers.p1]        # 定义了名字为p1的producer
    topic="sre-infra-test"  # 指定生产消息的topic
    balancer="my-balancer"  # 指定balancer，此balancer非默认balancer，需要使用ekafka.WithRegisterBalancer()注册
[kafka.consumers.c1]        # 定义了名字为c1的consumer
    topic="sre-infra-test"  # 指定消费的topic
    groupID="group-1"       # 如果配置了groupID，将初始化为consumerGroup	
[kafka.consumers.c2]        # 定义了名字为c2的consumer
    topic="sre-infra-test"  # 指定消费的topic
    groupID="group-2"       # 如果配置了groupID，将初始化为consumerGroup	
```

## 2.4 用户代码
```go
package main
// produce 生产消息
func produce(w *ekafka.Producer) {
    // 生产3条消息
    err := w.WriteMessages(context.Background(),
        ekafka.Message{Key: []byte("Key-A"), Value: []byte("Hello World!")},
        ekafka.Message{Key: []byte("Key-B"), Value: []byte("One!")},
        ekafka.Message{Key: []byte("Key-C"), Value: []byte("Two!")},
    )
    if err != nil {
        log.Fatal("failed to write messages:", err)
    }
    if err := w.Close(); err != nil {
        log.Fatal("failed to close writer:", err)
    }
    fmt.Println(`produce message succ--------------->`)
}
    
// consume 使用consumer/consumerGroup消费消息
func consume(r *ekafka.Consumer) {
    ctx := context.Background()
    for {
        // ReadMessage 再收到下一个Message时，会阻塞
        msg, err := r.ReadMessage(ctx)
        if err != nil {
            panic("could not read message " + err.Error())
        }
        // 打印消息
        fmt.Println("received: ", string(msg.Value))
        err = r.CommitMessages(ctx, msg)
        if err != nil {
            log.Printf("fail to commit msg:%v", err)
        }
    }
}

func main() {
    var stopCh = make(chan bool)
    // 假设你配置的toml如下所示
    conf := `
    [kafka]
        debug=true
        brokers=["localhost:9091","localhost:9092","localhost:9093"]
        [kafka.client]
            timeout="3s"
        [kafka.producers.p1]        # 定义了名字为p1的producer
            topic="sre-infra-test"  # 指定生产消息的topic
            balancer="my-balancer"  # 指定balancer，此balancer非默认balancer，需要使用ekafka.WithRegisterBalancer()注册
        [kafka.consumers.c1]        # 定义了名字为c1的consumer
            topic="sre-infra-test"  # 指定消费的topic
            groupID="group-1"       # 如果配置了groupID，将初始化为consumerGroup	
        [kafka.consumers.c2]        # 定义了名字为c2的consumer
            topic="sre-infra-test"  # 指定消费的topic
            groupID="group-2"       # 如果配置了groupID，将初始化为consumerGroup	
    `
    // 加载配置文件
    err := econf.LoadFromReader(strings.NewReader(conf), toml.Unmarshal)
    if err != nil {
    panic("LoadFromReader fail," + err.Error())
    }
    
    // 初始化ekafka组件
    cmp := ekafka.Load("kafka").Build(
    // 注册名为my-balancer的自定义balancer
    ekafka.WithRegisterBalancer("my-balancer", &kafka.Hash{}),
    )
    
    // 使用p1生产者生产消息
    go produce(cmp.Producer("p1"))
    
    // 使用c1消费者消费消息
    consume(cmp.Consumer("c1"))
    
    stopCh <- true
}
```

## 3 Consumer
### 3.1 Consumer配置
```go
type config struct {
    // Brokers brokers地址
    Brokers []string `json:"brokers" toml:"brokers"`
    // Debug 是否开启debug模式
    Debug bool `json:"debug" toml:"debug"`
    // Client 用于创建topic等
    Client clientConfig `json:"client" toml:"client"`
    // Producers 多个消费者，用于生产消息
    Producers map[string]producerConfig `json:"producers" toml:"producers"`
    // Consumers 多个生产者，用于消费消息
    Consumers    map[string]consumerConfig `json:"consumers" toml:"consumers"`
}

type consumerConfig struct {
    // Partition 指定分区ID，和GroupID不能同时配置
    Partition int `json:"partition" toml:"partition"`
    // GroupID 指定分组ID，和Partition不能同时配置，当配置了GroupID时，默认使用ConsumerGroup来消费
    GroupID string `json:"groupID" toml:"groupID"`
    // Topic 消费的topic
    Topic string `json:"topic" toml:"topic"`
    // MinBytes 向kafka发送请求的包最小值
    MinBytes int `json:"minBytes" toml:"minBytes"`
    // MaxBytes 向kafka发送请求的包最大值
    MaxBytes int `json:"maxBytes" toml:"maxBytes"`
    // WatchPartitionChanges 是否监听分区变化
    WatchPartitionChanges bool `json:"watchPartitionChanges" toml:"watchPartitionChanges"`
    // PartitionWatchInterval 监听分区变化时间周期
    PartitionWatchInterval time.Duration `json:"partitionWatchInterval" toml:"partitionWatchInterval"`
    // RebalanceTimeout rebalance 超时时间
    RebalanceTimeout time.Duration `json:"rebalanceTimeout" toml:"rebalanceTimeout"`
    // MaxWait 从kafka批量获取数据时，最大等待间隔
    MaxWait time.Duration `json:"maxWait" toml:"maxWait"`
    // ReadLagInterval 获取消费者滞后值的时间周期
    ReadLagInterval   time.Duration `json:"readLagInterval" toml:"readLagInterval"`
    HeartbeatInterval time.Duration `json:"heartbeatInterval" ,toml:"heartbeatInterval"`
    CommitInterval    time.Duration `json:"commitInterval" toml:"commitInterval"`
    SessionTimeout    time.Duration `json:"sessionTimeout" toml:"sessionTimeout"`
    JoinGroupBackoff  time.Duration `json:"joinGroupBackoff" toml:"joinGroupBackoff"`
    RetentionTime     time.Duration `json:"retentionTime" toml:"retentionTime"`
    StartOffset       int64         `json:"startOffset" toml:"startOffset"`
    ReadBackoffMin    time.Duration `json:"readBackoffMin" toml:"readBackoffMin"`
    ReadBackoffMax    time.Duration `json:"readBackoffMax" toml:"readBackoffMax"`
}
```
### 3.2 用户配置
```toml
	[kafka]
        debug=true
        brokers=["localhost:9094"]
	[kafka.client]
        timeout="3s"
	[kafka.producers.p1]        # 定义了名字为p1的producer
		topic="sre-infra-test"  # 指定生产消息的topic
	[kafka.consumers.c1]        # 定义了名字为c1的consumer
		topic="sre-infra-test"  # 指定消费的topic
		groupID="group-1"       # 如果配置了groupID，将初始化为consumerGroup	
	[kafkaConsumerServers.s1]
        debug=true
        consumerName="c1"
```
### 3.3 用户代码
```go
package main

func main() {
	conf := `
	[kafka]
	debug=true
	brokers=["localhost:9094"]
	[kafka.client]
        timeout="3s"
	[kafka.producers.p1]        # 定义了名字为p1的producer
		topic="sre-infra-test"  # 指定生产消息的topic

	[kafka.consumers.c1]        # 定义了名字为c1的consumer
		topic="sre-infra-test"  # 指定消费的topic
		groupID="group-1"       # 如果配置了groupID，将初始化为consumerGroup	

	[kafkaConsumerServers.s1]
	debug=true
	consumerName="c1"
`
	// 加载配置文件
	err := econf.LoadFromReader(strings.NewReader(conf), toml.Unmarshal)
	if err != nil {
		panic("LoadFromReader fail," + err.Error())
	}

	app := ego.New().Serve(
		// 可以搭配其他服务模块一起使用
		egovernor.Load("server.governor").Build(),

		// 初始化 Consumer Server
		func() *consumerserver.Component {
			// 依赖 `ekafka` 管理 Kafka consumer
			ec := ekafka.Load("kafka").Build()
			cs := consumerserver.Load("kafkaConsumerServers.s1").Build(
				consumerserver.WithEkafka(ec),
			)

			// 用来接收、处理 `kafka-go` 和处理消息的回调产生的错误
			consumptionErrors := make(chan error)

			// 注册处理消息的回调函数
			cs.OnEachMessage(consumptionErrors, func(ctx context.Context, message kafka.Message) error {
				fmt.Printf("got a message: %s\n", string(message.Value))
				// 如果返回错误则会被转发给 `consumptionErrors`
				return nil
			})

			return cs
		}(),
		// 还可以启动多个 Consumer Server
	)
	if err := app.Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}
```