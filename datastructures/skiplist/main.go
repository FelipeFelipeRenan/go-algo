package main

import (
	"fmt"
	"math/rand"
)

const maxLevel = 16
const p = 0.5

type Node struct {
	Value string

	Forward []*Node
}

type SkipList struct {
	Header *Node
	Level int
}


func NewSkipList() *SkipList{

	header := &Node{Forward: make([]*Node, maxLevel) }
	return &SkipList{Header: header, Level: 1}
}

func (sl *SkipList) randomLevel() int  {
	level := 1

	for rand.Float64() < p && level < maxLevel{
		level++
	}
	return level
}


func (sl *SkipList) Insert(value int)  {
	
	fmt.Printf("Inserindo '%d'...\n", value)

	update := make([]*Node, maxLevel)
	currentNode := sl.Header

	
}