package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello Channel!"
	}()

	message := <- channel
	fmt.Println(message)
}