package main

func reverseLinkedList(node *Node) *Node {
	var prev *Node
	for node != nil {
		aux := node.Next
		node.Next = prev
		prev = node
		node = aux
	}
	return prev
}
