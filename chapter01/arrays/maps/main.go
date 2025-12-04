package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["One"] = 1
	myMap["Two"] = 2
	myMap["Three"] = 3

	//fmt.Println(myMap, myMap["One"])

	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello": "golang"}`)
	err := json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", myJsonMap["hello"])
}