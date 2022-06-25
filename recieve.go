package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	var m = make(map[string]string)
	conn, err := amqp.Dial("amqp://guest:guest@localhost/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")
	content, con, err := ch.Get("hello", true)
	if err != nil {
		fmt.Println(con)
	}
	err1 := json.Unmarshal(content.Body, &m)
	if err1 != nil {
		fmt.Println("error")
	}
	fmt.Println(m["omidgharahkol"])

}
