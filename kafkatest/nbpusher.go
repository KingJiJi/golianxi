package main

import (
	"dbms/lib/tokenbucket"
	"flag"
	"fmt"
	"os"
	"time"

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
	flag.IntVar(&partition, "nn", 20000, "msgs per  second")
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
	rangnizou := tokenbucket.NewBucket(int64(partition), 1*time.Second)
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
	myticker := time.NewTicker(5 * time.Second)

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
		olddata := 0
		for {
			select {

			case <-myticker.C:
				fmt.Printf("total published msg= %d\n", data)
				fmt.Printf("published msg= %d  in 5 seconds\n", data-olddata)
				olddata = data

			default:
				if o := rangnizou.Take(1); o != 0 {
					data++
					pMessage := &sarama.ProducerMessage{
						Topic: topic,
						Key:   nil,
						Value: sarama.ByteEncoder([]byte(message)),
					}

					pInChan <- pMessage
				}
			}
		}
	}
	producer.Close()
	client.Close()

}
