package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main()  {

	log.Println("RabbitMQ test example")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		panic(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err !=nil {
		log.Println(err)
	}

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte("Hello, seongwon"),
		})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully Published Message to Queue")

}
