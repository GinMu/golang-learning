package main

import "fmt"

// Node define list node
type Node struct {
	Data int
	Next *Node
}

func showNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func main() {
	head := new(Node)
	head.Data = 1
	node1 := new(Node)
	node1.Data = 2
	head.Next = node1
	node2 := new(Node)
	node2.Data = 3
	node1.Next = node2
	showNode(head)
}
