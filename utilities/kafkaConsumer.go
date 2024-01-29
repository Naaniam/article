package utilities

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartKafkaConsumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "article-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		return
	}

	err = consumer.SubscribeTopics([]string{"article-topic"}, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if kafkaError, ok := err.(kafka.Error); ok && kafkaError.IsFatal() {
			log.Printf("Fatal consumer error: %v (%v)\n", err, msg)
			return
		} else {
			log.Printf("Non-fatal consumer error: %v (%v)\n", err, msg)
		}
	}

}

