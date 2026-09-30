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

	queueName := "simple"

	// 声明队列
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久性
		false,     // 自动删除
		false,     // 排他性
		false,     // 非阻塞
		nil,       // 其他参数
	)
	common.FailOnError(err, "无法声明队列")

	// 接收消息
	msgs, err := ch.Consume(
		q.Name, // 队列
		"",     // 消费者
		false,  // 自动确认
		false,  // 排他性
		false,  // 非本地
		false,  // 非阻塞
		nil,    // 其他参数
	)
	common.FailOnError(err, "无法注册消费者")

	go func() {
		for d := range msgs {
			log.Printf("收到消息: %s\n", d.Body)
			// log.Println("回复消息")
			// d.Ack(false)
		}
	}()
	log.Println("等待接收消息")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Printf("已关闭")
}
