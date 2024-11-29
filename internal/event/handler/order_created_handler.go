package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/IsJordanBraz/go-clean-architecture/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(RabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: RabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order Createed: %v", event.GetPayload())

	jsonOutput, _ := json.Marshal(event.GetPayload())

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", //exchange
		"",           //key name
		false,        //mandatory
		false,        //immediate
		message,      //message to publish
	)
}
