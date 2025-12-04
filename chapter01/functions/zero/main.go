package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err := divisilbleBy(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", res)
}

func divisilbleBy(n, divisor int) (bool, error) {
	if divisor == 0 {
		return false, errors.New("A number cannot be divided by zero")
	}

	return (n % divisor == 0), nil
}