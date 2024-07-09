package main

import (
	"AdapterPattern/adapter"
	"AdapterPattern/services"
)

func main() {
	//first initial requirement
	apple := services.Apple{}
	client := services.Client{}
	client.ChargeMobile(&apple)

	//extended requirement
	android := &services.Android{}
	androidAdapter := &adapter.AndroidAdapter{
		Android: android,
	}
	androidAdapter.ChargeAppleMobile()
}
