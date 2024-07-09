package services

import "fmt"

type On struct {
}

func (o *On) State() {
	fmt.Println("Tv is on")
}
