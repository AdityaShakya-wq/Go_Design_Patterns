package main

import "ObserverDesignPattern/services"

func main() {
	pub := services.GetPublisher()

	sub := services.GetSubscriber("1")
	sub1 := services.GetSubscriber("2")
	sub2 := services.GetSubscriber("3")
	pub.Register(sub)
	pub.Register(sub1)
	pub.Register(sub2)
	pub.NotifyAll("Notification")
}
