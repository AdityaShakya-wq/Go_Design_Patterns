package services

import "fmt"

// proto type
type Apple struct {
}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("Charge apple mobile!")
}
