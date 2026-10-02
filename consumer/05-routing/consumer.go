package main

import (
	"log"
	"mq/common"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := common.InitMQ()
	defer ch.Close()

	// 声明交换器
	exChangeName := common.ExchangeLogsDirect(ch)

	// 创建临时队列（自动生成唯一名称，连接关闭后自动删除），从而只消费新消息
	q, err := ch.QueueDeclare(
		"",    // 空名称，让 RabbitMQ 自动生成
		false, // 非持久化
		true,  // 排他性（仅当前连接可见）
		true,  // 自动删除，连接关闭后队列自动删除
		false, // 不等待服务器确认
		nil,   // 额外参数
	)
	if err != nil {
		log.Fatalf("声明临时队列失败 %s", err)
	}

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("绑定队列 %s 到 %s 交换器，key： %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key 订阅交换器中 key 为 s 的消息
			exChangeName, // exchange
			false,
			nil)
		common.FailOnError(err, "Failed to bind a queue")
	}

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者标签
		true,   // 自动确认（简单示例，生产环境建议手动确认）
		false,  // 非排他性
		false,  // 非本地消费者
		false,  // 不等待服务器确认
		nil,    // 额外参数
	)
	if err != nil {
		log.Fatalf("创建消费者失败 %s", err)
	}

	log.Println("准备接收消息")

	go func() {
		for d := range msgs {
			log.Printf("[x] 收到消息: %s", d.Body)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Printf("已关闭")
}
