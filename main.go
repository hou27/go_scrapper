package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines
	go parallelCount("hou27")
	go parallelCount("jalapeno")
	time.Sleep(time.Second * 3)
}

func parallelCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}