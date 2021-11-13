package main

import (
	"fmt"
	"time"
)

func main() {
	// channel
	c := make(chan bool)
	people := [2]string{"jalapeno", "hou27"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c) // main func wait until reply
	fmt.Println(<-c)
	// fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	c <- true // send true to channel
}