package main

import "time"

func main() {
	messagePrinter := func(msg string) {
		println(msg)
	}

	go messagePrinter("Design Pattern")
	go messagePrinter("goroutine")

	time.Sleep(time.Second)
}