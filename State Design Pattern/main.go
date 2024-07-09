package main

import "StateDesignPattern/services"

// client
func main() {
	tvContext := services.GetState() //default state is off
	tvContext.GetState()
	tvContext.SetState(&services.On{})
	tvContext.GetState()
}
