package main

import "fmt"

func main() {
	mySlice := make([]int, 10)

	mySlice = append(mySlice, 6)

	mySlice = mySlice[1:]

	mySlice = append(mySlice[1:], mySlice[2:]...)

	fmt.Println(mySlice)
}