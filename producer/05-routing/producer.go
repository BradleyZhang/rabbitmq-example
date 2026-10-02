package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"mq/common"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := common.InitMQ()
	defer ch.Close()

	exchangeName := common.ExchangeLogsDirect(ch)

	// body := bodyFrom(os.Args)

	reader := bufio.NewReader(os.Stdin)
	routingKey := severityFrom(os.Args)
	fmt.Println("即将发送 key 为：" + routingKey + " 的消息")
	fmt.Println("输入消息")
	for {
		fmt.Print("> ")
		body, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatal(err.Error())
			}
		}
		err = ch.Publish(
			exchangeName,
			routingKey, // routing key 发送到交换器时带key
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		common.FailOnError(err, "发送失败")
	}

}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
