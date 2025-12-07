package main

import "fmt"

type Tree struct {
	LeafValue int
	Right *Tree
	Left *Tree
}

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p *Parent) int {
	fmt.Println(p.SomeField)
	return p.SomeField
}

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right: &Tree{6, nil, nil},
			Left: nil,
		},
		Left: &Tree{4, nil, nil},
	}

	son := Son{}
	GetParentField(&son.P)

	fmt.Println(root.Right.Right.LeafValue)
}