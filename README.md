RabbitMQ 学习用示例代码

## Run
## 安全配置
修改 ca.sh 中 IP.1 为实际地址
```bash
mkdir ca
cd ca
../ca.sh
```

启动 RabbitMQ
### Docker
```bash
docker compose up -d
```
控制台 `https://127.0.0.1:15671`
账号 admin
密码 password

### 01 简单模式
生产者 
```bash
make p m=message
```

消费者
```bash
make c
```

内容：连接RabbitMQ、创建通道、声明队列
-接收消息、确认消息
-发送消息
#### 02 安全模式
生产者 
```bash
make sp m=message
```

消费者
```bash
make sc
```

内容：配置TLS
### 03 工作队列
生产者
```bash
$ make wp m="process for 3 seconds ..."
```

消费者
```bash
$ make wc
```
内容：持久化：队列持久化 消息持久化、Qos公平分发
### 04 发布订阅模式
生产者
```bash
$ make 4p
```

消费者
```bash
$ make 4c
```

内容：声明交换器、临时队列、队列交换器绑定

### 05 路由：队列订阅交换器消息的子集
生产者
```bash
$ make 5p key=info
$ make 5p key=warning
$ make 5p key=error
```

消费者
```bash
make 5c key="warning info"
make 5c key="error warning
```

内容：direct 类型交换器，key