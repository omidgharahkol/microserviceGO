package MessageBrocker

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

func Receive(queue string) map[string]string {

	var OrderDetail = make(map[string]string)
	conn, err := amqp.Dial("amqp://guest:guest@localhost/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")
	content, con, err := ch.Get(queue, true)
	if err != nil {
		fmt.Println(con)
	}
	err1 := json.Unmarshal(content.Body, &OrderDetail)
	if err1 != nil {
		fmt.Println(err1)
	}
	return OrderDetail

}

func H() string {
	return "omid"
}
