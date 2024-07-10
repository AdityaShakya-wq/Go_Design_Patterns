package main

import (
	"fmt"
	"sync"
)

type config struct {
}

var counter int = 1
var singleConfigInstance *config
var mutex = sync.Mutex{}

func getConfigInstance() *config {
	if singleConfigInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creating single instance now, and counter is ", counter)
			singleConfigInstance = &config{}
			counter++
		} else {
			fmt.Println("Single instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single instance already created-2, returning that same")
	}

	return singleConfigInstance
}

func main() {
	for i := 1; i < 50; i++ {
		go getConfigInstance()
	}
	fmt.Scanln()
}
