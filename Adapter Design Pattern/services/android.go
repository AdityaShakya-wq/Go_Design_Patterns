package services

import "fmt"

type Android struct{}

func (a *Android) ChargeAndroidMobile() {
	fmt.Println("Charge android mobile")
}
