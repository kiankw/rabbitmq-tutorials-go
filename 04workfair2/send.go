package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

// 工作队列 公平调度模式 Fair dispatch

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接工厂
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建信道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 绑定队列
	q, err := ch.QueueDeclare(
		"work_fair", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for i := 0; i < 20; i++ {
		body := "Hello World!" + strconv.Itoa(i)
		// 发送消息
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		fmt.Println(" [x] Sent '"+body+"'", i)
		failOnError(err, "Failed to publish a message")
	}
}
