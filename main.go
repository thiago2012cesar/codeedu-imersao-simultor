package main

import (
	"fmt"
	kafka2 "github.com/thiago2012cesar/codeedu-imersao-simultor/application/kafka"
	"github.com/thiago2012cesar/codeedu-imersao-simultor/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
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