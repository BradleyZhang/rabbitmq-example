package main

import (
	"log"
	"mq/common"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := common.InitMQ()
	defer ch.Close()

	// 声明队列
	q, err := common.QueueWork(ch)
	common.FailOnError(err, "无法声明队列")
	// 发送消息
	body := getMsg()
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息持久化
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	common.FailOnError(err, "发送消息失败")
	log.Printf("发送消息: %s\n", body)
}
func getMsg() string {
	msg := "default msg"
	args := os.Args
	if len(args) > 1 && args[1] != "" {
		msg = strings.Join(args[1:], " ")
	}
	return msg
}
