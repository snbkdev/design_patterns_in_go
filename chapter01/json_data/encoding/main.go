package main

import (
	"encoding/json"
	"fmt"
)

type myObject struct {
	//Number int `json:"number"`
	//Number int
	//Word string

	Number int `json:"number"`
	Word string `json:"string"`
}

func main() {
	// object := myObject{5, "Packt"}
	// ojson, _ := json.Marshal(object)
	// fmt.Printf("%s\n", ojson)

	// josonBytes := []byte(`{"number":5, "string": "packt"}`)
	// var object myObject
	// err := json.Unmarshal(josonBytes, &object)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("NUmber is %d, Word is %s\n", object.Number, object.Word)

	jsonBytes := []byte(`{"number": 5, "string": "packt"}`)
	var dangerousObject map[string]interface{}
	err := json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number is %d, ", dangerousObject["number"])
	fmt.Printf("Word is %s\n", dangerousObject["string"])
	fmt.Printf("Error reference is %v\n", dangerousObject["nothing"])
}