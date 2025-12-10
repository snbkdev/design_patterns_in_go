package main

import "time"

func main() {
	go helloWorld2()

	time.Sleep(time.Second)
}

func helloWorld2() {
	println("Hello!!!")
}