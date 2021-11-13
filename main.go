package main

import (
	"fmt"
	"time"
)

func main() {
	// channel
	c := make(chan string)
	people := [2]string{"jalapeno", "hou27"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c) // main func wait until reply(receiving is blocking operation)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy" // send to channel
}