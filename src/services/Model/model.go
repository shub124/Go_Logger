package model

import (
	"fmt"
	meta "services/Common/Meta"
	libserver "services/Libraries/Server"

	sarama "github.com/Shopify/sarama"
)

func Cricketlog(data *meta.Cricketmeta, RemoteAddr string) {
	fmt.Println(data)
	libserver.Iserver.CricketLog.Input() <- &sarama.ProducerMessage{
		Topic: "Cricket_Log",
		Key:   sarama.StringEncoder(RemoteAddr),
		Value: data,
	}

}

func Userlog(data *meta.Usermeta, RemoteAddr string) {
	libserver.Iserver.UserLog.Input() <- &sarama.ProducerMessage{
		Topic: "User_Log",
		Key:   sarama.StringEncoder(RemoteAddr),
		Value: data,
	}
}
