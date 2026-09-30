package main

import (
	"fmt"
	"log"
	"mq/common"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
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
	if err != nil {
		log.Fatalf("无法声明队列: %v", err)
	}

	// 发送消息
	body := getMsg()

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "发送失败: %v", err)
		return
	}
	fmt.Printf("发送消息: %s\n", body)
}
func getMsg() string {
	msg := "default msg"
	args := os.Args
	if len(args) > 1 {
		msg = args[1]
	}
	return msg
}
