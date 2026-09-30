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

### 简单模式
生产者 
```bash
make p m=message
```

消费者
```bash
make c
```

#### 安全模式
生产者 
```bash
make sp m=message
```

消费者
```bash
make sc
```
### 发布订阅模式