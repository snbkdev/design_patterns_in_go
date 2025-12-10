package main

import "time"

func main() {
	go func(msg string) {
		println(msg)
	} ("Design Patterns!")
	time.Sleep(time.Second)
} 