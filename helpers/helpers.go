package helpers

import "github.com/confluentinc/confluent-kafka-go/kafka"

const (
	KafkaBootstrapServers = "localhost:9092"
	KafkaTopic            = "article-topic"
	GroupID               = "article-consumer-group"
)

// Kafka Producer
var KafkaProducer *kafka.Producer

// Kafka consumer
var KafkaConsumer *kafka.Consumer
