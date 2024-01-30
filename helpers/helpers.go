package helpers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

const (
	KafkaBootstrapServers = "localhost:9092"
	KafkaTopic            = "article-topic"
	GroupID               = "article-consumer-group"
)

// Kafka Producer
var KafkaProducer *kafka.Producer

// Kafka consumer
var KafkaConsumer *kafka.Consumer

// Log
var Log *logrus.Logger
