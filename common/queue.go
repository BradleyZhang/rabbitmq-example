package common

import amqp "github.com/rabbitmq/amqp091-go"

func QueueWork(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"work", // 队列名称
		true,   // 队列持久性
		false,  // 未使用后删除
		false,  // 排他性
		false,  // 非阻塞
		nil)    // 其他参数
	return &q, err
}
