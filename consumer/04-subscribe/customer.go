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

	exChangeName := common.ExchangeLogs(ch)

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

	// 将队列绑定到交换器（关键步骤）
	err = ch.QueueBind(
		q.Name,       // 队列名称
		"",           // 路由键（fanout 类型忽略）
		exChangeName, // 交换器名称
		false,        // 不等待服务器确认
		nil,          // 额外参数
	)
	common.FailOnError(err, "队列绑定交换器失败")

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
