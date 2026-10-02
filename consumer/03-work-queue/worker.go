package main

import (
	"bytes"
	"log"
	"mq/common"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch := common.InitMQ()
	defer ch.Close()

	// 声明队列
	q, err := common.QueueWork(ch)
	common.FailOnError(err, "无法声明队列")

	err = ch.Qos(
		1, // 预取计数为1,在确认了1条消息后，RabbitMQ才分发新消息
		0,
		false,
	)
	common.FailOnError(err, "Oos 设置失败")
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
			dotCount := bytes.Count(d.Body, []byte(".")) // 为消息中每个点模拟一秒工作
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Println("完成")
			d.Ack(false)
			// 调试忘记 Ack 的错误： `sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged`
		}
	}()
	log.Println("等待接收消息")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Printf("已关闭")
}
