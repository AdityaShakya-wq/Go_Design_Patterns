package services

import "fmt"

type ToSubscribe interface {
	ReactToPublish(msg string)
}

type Subscriber struct {
	SubscriberId string
}

func GetSubscriber(id string) Subscriber {
	return Subscriber{SubscriberId: id}
}

func (s Subscriber) ReactToPublish(msg string) {
	fmt.Println("Subscriber received message", msg, " for subscriber", s.SubscriberId)
}
