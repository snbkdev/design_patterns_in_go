package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicit string = "Hello, it is a variable"

	inferred := "second variable"

	fmt.Println("Variable explicit is of type:", reflect.TypeOf(explicit))
	fmt.Println("Variable inferred is of type:", reflect.TypeOf(inferred))
}
