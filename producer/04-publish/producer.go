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

	exchangeName := common.ExchangeLogs(ch)

	// body := bodyFrom(os.Args)

	reader := bufio.NewReader(os.Stdin)
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
		err = ch.Publish(exchangeName, "", false, false, amqp091.Publishing{ContentType: "text/plain", Body: []byte(body)})
		common.FailOnError(err, "发送失败")
	}

}
