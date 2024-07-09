package services

import "fmt"

type ToPublish interface {
	Register(s Subscriber)
	NotifyAll(msg string)
}

type Publisher struct {
	Subs []ToSubscribe
}

func GetPublisher() Publisher {
	return Publisher{
		Subs: make([]ToSubscribe, 0),
	}
}

func (p *Publisher) Register(s ToSubscribe) {
	p.Subs = append(p.Subs, s)
}

func (p *Publisher) NotifyAll(msg string) {
	for _, subs := range p.Subs {
		fmt.Println("Publisher notifying subscriber with id", subs.(Subscriber).SubscriberId)
		subs.ReactToPublish(msg)
	}
}
