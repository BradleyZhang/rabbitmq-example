package common

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// 声明 exchange
func ExchangeLogs(ch *amqp.Channel) (name string) {
	name = "logs"
	err := ch.ExchangeDeclare(
		name,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "exchange 声明失败")
	return name
}
