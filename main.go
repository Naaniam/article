package main

import (

	// user defined packages
	"article/drivers"
	"article/helpers"
	"article/migrations"
	"article/repository"
	"article/router"
	"article/utilities"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("error occured when getting hostname:", err)
		return
	}
	// Initialize Kafka producer
	helpers.KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": helpers.KafkaBootstrapServers,
		"client.id":         hostName,
		"acks":              "all",
	})
	if err != nil {
		return
	}
	fmt.Println("Kafka Producer", helpers.KafkaProducer)
}

func main() {
	go utilities.StartKafkaConsumer()
	// Establishing a DB-connection
	connection := drivers.SQLDriver()
	// Checking for Database updates
	migrations.Migrations(connection)

	// Routing all handlers
	router.Routing(repository.NewDbConnection(connection))
	//go utilities.StartKafkaConsumer()

}
