package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/gabrielsos/fullcycle-simulator-api/app/kafka"
	"github.com/gabrielsos/fullcycle-simulator-api/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("error loading .env file" + error.Error())
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
