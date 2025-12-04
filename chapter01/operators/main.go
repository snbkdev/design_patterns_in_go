package main

import "fmt"

func main() {
	if "a" == "b" || 10 == 10 || true == false {
		println("10 is equal to 10")
	} else if 11 == 11 && "go" == "go" {
		println("This is not print because previous condition was satisfied")
	} else {
		println("In casde no condition is satisfied, print this")
	}

	number := 3
	switch (number) {
	case 1:
		println("Number is", number)
	case 2:
		println("Number is", number)
	case 3:
		println("Number is", number)
	}

	for i := 0; i <= 10; i++ {
		println(i)
	}

	slice1 := []int{10, 20, 30, 40, 50}

	for index, value := range slice1 {
		fmt.Printf("Index is %d nad value is %d", index, value)
	}

	for index := 0; index < len(slice1); index++ {
		value := slice1[index]
		fmt.Printf("Index is %d and value is %d", index, value)
	}

}