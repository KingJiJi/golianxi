package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

var hostname string
var topic string
var group string

func init() {
	flag.StringVar(&hostname, "hostname", "localhost:9092", "host:port string for the kafka server")
	flag.StringVar(&topic, "topic", "test", "topic to publish to")
	flag.StringVar(&group, "group", "", "consumer group")
}
func main() {
	flag.Parse()
	fmt.Println(" ---------------------- ")
	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	//config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false
	//config.Group.Return.Notifications = true

	//sarama config
	config.Config.Consumer.Fetch.Default = 32768 //10240 //default 32768  32k
	//config.Config.Consumer.Fetch.Max = 0
	config.Config.ChannelBufferSize = 1024

	cg, err := cluster.NewConsumer(strings.Split(hostname, ","), group, []string{topic}, config)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	myticker := time.NewTicker(5 * time.Second)
	var precount int64 = 0
	var nowcount int64 = 0
	//idledur := time.Duration(2 * time.Second)
	for {
		select {
		case ev, ok := <-cg.Messages():
			if ok {
				cg.MarkOffset(ev, "")
				nowcount++
			}

		case ntf, more := <-cg.Notifications():
			if more {
				fmt.Printf("kafkaoutput Rebalanced: %+v\n", ntf)
			}
		case err, more := <-cg.Errors():
			if more {
				fmt.Printf("kafkaoutput Error: %s\n", err.Error())
			}

		case <-myticker.C:
			fmt.Printf("kafkaconsumer consumer %s:%d msgs in 5 second\n", topic, nowcount-precount)
			precount = nowcount
		}
	}

	cg.Close()

}
