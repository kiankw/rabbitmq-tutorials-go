package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalln("%s: %s", msg, err)
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

	// 绑定交换机
	err = ch.ExchangeDeclare(
		"exchange_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	message1 := "Hello World! 1"
	message2 := "Hello World! 2"
	message3 := "Hello World! 3"
	routingKey1 := "quick.orange.rabbit"
	routingKey2 := "lazy.pink.rabbit"
	routingKey3 := "quick.orange.male.rabbit"

	err = ch.Publish(
		"exchange_topic",
		routingKey1,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message1),
		})
	log.Println(" [x] Sent '" + message1 + "'")
	failOnError(err, "Failed to publish a message")

	err = ch.Publish(
		"exchange_topic",
		routingKey2,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message2),
		})
	log.Println(" [x] Sent '" + message2 + "'")
	failOnError(err, "Failed to publish a message")

	err = ch.Publish(
		"exchange_topic",
		routingKey3,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message3),
		})
	log.Println(" [x] Sent '" + message3 + "'")
	failOnError(err, "Failed to publish a message")
}
