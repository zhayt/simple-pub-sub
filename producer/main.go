package main

import (
	"encoding/json"
	"flag"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"log"
	"math/rand"
)

var (
	bootstrapServer = flag.String("bootstrap-server", "localhost:9093", "Kafka bootstrap server to connect to")
	topic           = flag.String("topic", "log", "Kafka topic name")
)

type Message struct {
	UUID   string `json:"uuid"`
	Action string `json:"action"`
}

func main() {
	flag.Parse()

	producer, err := sarama.NewSyncProducer([]string{*bootstrapServer}, nil)
	if err != nil {
		log.Fatalf("Init new producer error: %s", err.Error())
	}

	actions := []string{"click", "failed", "error", "success", "save", "reload"}
	for i := 0; i < 100; i++ {
		message := &Message{UUID: uuid.NewString(), Action: actions[rand.Intn(len(actions)-1)]}

		data, err := json.Marshal(message)
		if err != nil {
			log.Fatalf("Marshal data error: %s", err.Error())
		}

		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(data),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("Send message error: %s", err.Error())
		}
		log.Printf("Message sent! [Partiton: %d][Offset: %d]\n", partition, offset)
	}

	log.Println("END")
}
