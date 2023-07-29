package event

import amqp "github.com/rabbitmq/amqp091-go"

func declareExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		"logs_topic", // Name of the exchange
		"topic",	  // Type
		true,			// Durability of the exchange
		false,			// Auto-deleted?
		false,			// internal communication OR between microservices
		false,			// No-wait?
		nil,			// No specific arguments
	)
}

func declareRandomQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		"", // name
		false, // durable?
		false, // dont delete when it is unused
		true, // exclusive (dont share it around)
		false, // no-wait?
		nil, //no specific arguments
	)
}