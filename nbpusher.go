package main

import (
	"flag"
	"fmt"
	"os"

	kafka "github.com/jdamick/kafka"
	"gopkg.in/Shopify/sarama.v1"
)

var hostname string
var topic string
var partition int
var message string
var messageFile string
var compress bool

func init() {
	flag.StringVar(&hostname, "hostname", "localhost:9092", "host:port string for the kafka server")
	flag.StringVar(&topic, "topic", "test", "topic to publish to")
	flag.IntVar(&partition, "partition", 0, "partition to publish to")
	flag.StringVar(&message, "message", "", "message to publish")
	flag.StringVar(&messageFile, "messagefile", "", "read message from this file")
	flag.BoolVar(&compress, "compress", false, "compress the messages published")
}

func processKafkaErrors(errChan <-chan *sarama.ProducerError) {

	var (
		ok   = true
		pErr *sarama.ProducerError
	)
	for ok {
		select {
		case pErr, ok = <-errChan:
			if !ok {
				break
			}
			err := pErr.Err
			switch err.(type) {
			case sarama.PacketEncodingError:
				fmt.Println("Error: ", err)

			default:
				if err != nil {
					fmt.Println("Error: ", err)
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	fmt.Println("Publishing :", message)
	fmt.Printf("To: %s, topic: %s, partition: %d\n", hostname, topic, partition)
	fmt.Println(" ---------------------- ")
	broker := kafka.NewBrokerPublisher(hostname, topic, partition)

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.NoResponse
	client, err := sarama.NewClient([]string{hostname}, saramaConfig)
	if err != nil {

		fmt.Println("Error: ", err)
		return
	}
	producer, err := sarama.NewAsyncProducer([]string{hostname}, saramaConfig)
	if err != nil {

		fmt.Println("Error: ", err)
		return
	}

	errChan := producer.Errors()
	pInChan := producer.Input()
	go processKafkaErrors(errChan)

	if len(message) == 0 && len(messageFile) != 0 {
		file, err := os.Open(messageFile)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		stat, err := file.Stat()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		payload := make([]byte, stat.Size())
		file.Read(payload)
		timing := kafka.StartTiming("Sending")

		if compress {
			broker.Publish(kafka.NewCompressedMessage(payload))
		} else {
			broker.Publish(kafka.NewMessage(payload))
		}

		timing.Print()
		file.Close()
	} else {

		data := 0
		for {

			data++
			pMessage := &sarama.ProducerMessage{
				Topic: topic,
				Key:   nil,
				Value: sarama.ByteEncoder([]byte(message)),
			}

			pInChan <- pMessage
			fmt.Printf("======================== %d\n", data)
		}
	}
	producer.Close()
	client.Close()

}
