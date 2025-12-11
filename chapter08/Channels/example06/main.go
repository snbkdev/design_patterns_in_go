package main

import "time"

func sendString(ch chan <- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <- chan string, quitCh chan <- bool) {
	for {
		select {
		case msg := <- helloCh:
			println(msg)
		case msg := <- goodbyeCh:
			println(msg)
		case <- time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting!!!")
			quitCh <- true
			break
		}
	}
}

func main() {
	helloCH := make(chan string, 1)
	goodbyeCH := make(chan string, 1)
	quitCH := make(chan bool)
	go receiver(helloCH, goodbyeCH, quitCH)

	go sendString(helloCH, "hello!!!")

	time.Sleep(time.Second)

	go sendString(goodbyeCH, "GoodBye!!!")
	<- quitCH
}