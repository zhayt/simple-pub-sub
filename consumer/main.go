package main

import (
	"context"
	"flag"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	bootstrapServer = flag.String("bootstrap-server", "localhost:9093", "Kafka bootstrap server to connect to")
	topic           = flag.String("topic", "log", "Kafka topic name")
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		// Process the message as per your requirement here
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	flag.Parse()
	log.Println("Start consuming...")
	groupID := "consumer-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup([]string{*bootstrapServer}, groupID, config)
	if err != nil {
		log.Fatalf("Init new consumer error: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-ch
		cancel()
	}()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			err = consumerGroup.Consume(ctx, []string{*topic}, exampleConsumerGroupHandler{})
			if err != nil {
				log.Fatalf("consume msg err: %s", err)
			}
		}
	}

	log.Println("END")
}
