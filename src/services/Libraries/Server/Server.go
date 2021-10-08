package libserver

import (
	"log"
	"net/http"

	sarama "github.com/Shopify/sarama"
)

type Server struct {
	UserLog    sarama.AsyncProducer
	CricketLog sarama.AsyncProducer
}

var Iserver *Server

func Init(brokerlist []string) {

	Iserver = &Server{

		UserLog:    NewAsyncLogProducer(brokerlist),
		CricketLog: NewAsyncLogProducer(brokerlist),
	}

}

func Run(addr string, mux *http.ServeMux) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Listening for requests on %s...\n", addr)
	return httpServer.ListenAndServe()
}

func Close() error {
	if err := Iserver.UserLog.Close(); err != nil {
		log.Println("Failed to shut down userlog producer cleanly", err)
	}

	if err := Iserver.CricketLog.Close(); err != nil {
		log.Println("Failed to shut down  cricketlog producer cleanly", err)
	}

	return nil
}

func NewAsyncLogProducer(brokerlist []string) sarama.AsyncProducer {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokerlist, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}
