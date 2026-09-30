package main

import (
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
		false,     // 未使用后删除
		false,     // 排他性
		false,     // 非阻塞
		nil)       // 其他参数
	common.FailOnError(err, "无法声明队列")
	// 发送消息
	body := getMsg()
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	common.FailOnError(err, "发送消息失败")
	log.Printf("发送消息: %s\n", body)
}
func getMsg() string {
	msg := "default msg"
	args := os.Args
	if len(args) > 1 {
		msg = args[1]
	}
	return msg
}
