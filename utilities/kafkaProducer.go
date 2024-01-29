package utilities

import (
	"article/helpers"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func PublishToKafka(message string, eventType string, id uint) {
	// Prepare message data
	messageData := map[string]interface{}{
		"event_type": eventType,
		"entity_id":  id,
		"timestamp":  time.Now().UnixNano() / int64(time.Millisecond),
		"message":    message,
	}

	// Convert message data to JSON
	jsonData, err := json.Marshal(messageData)
	if err != nil {
		log.Printf("Error marshalling JSON data: %s", err)
		return
	}

	kafkaTopic := "article-topic"
	deliveryChan := make(chan kafka.Event, 1)

	err = helpers.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return
	}

	// Wait for delivery report
	e := <-deliveryChan
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", msg.TopicPartition.Error)
	} else {
		fmt.Printf("Produced message to topic %s, partition %d, offset %d: %s\n",
			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset, string(msg.Value))
	}
}
