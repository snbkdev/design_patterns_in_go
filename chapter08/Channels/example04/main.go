package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func() {
		channel <- "Hello Channel! 1"
		channel <- "Hello Channel! 2"
		println("Finishing!!!")
	}()

	time.Sleep(time.Second)

	message := <- channel
	fmt.Println(message)
}