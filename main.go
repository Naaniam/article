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
	"io"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
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
	helpers.Log = logrus.New()

	// Create a new log file or open an existing one
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		helpers.Log.Fatal("Failed to open log file: ", err)
	}

	// Set the log output to the file
	multiWriter := io.MultiWriter(os.Stdout, file)
	helpers.Log.SetOutput(multiWriter)

	go utilities.StartKafkaConsumer()

	// Establishing a DB-connection
	connection := drivers.SQLDriver()
	// Checking for Database updates
	migrations.Migrations(connection)

	// Routing all handlers
	router.Routing(repository.NewDbConnection(connection))

}
