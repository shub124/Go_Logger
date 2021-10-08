package main

import (
	"flag"
	"fmt"

	"os"
	"os/signal"
	"strings"
	"time"

	controller "consumeservice/Controller"

	"log"

	cluster "github.com/bsm/sarama-cluster"
)

func main() {
	brokers := flag.String("brokers", "", "brokers")
	//consumergroup := flag.String("consumergroup", "CONSUMER_GROUP", "Unique consumer group id")
	topics := flag.String("topics", "TOPIC_ID", "all topics")
	flag.Parse()

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = time.Second

	brokerlist := strings.Split(*brokers, ",")
	topiclist := strings.Split(*topics, ",")
	fmt.Println("Consumer service")
	consumer, err := cluster.NewConsumer(brokerlist, "my-consumer-group", topiclist, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "")
				controllerfunc, ok := controller.Controllermap[msg.Topic]
				if ok {
					controllerfunc(msg.Value)
				} // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
