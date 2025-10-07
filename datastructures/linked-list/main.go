package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) AddAtStart(value int) {
	newNode := &Node{Value: value}
	l.Head = newNode

	if l.IsEmpty() {
		l.Tail = newNode
	}
	l.Size++
	fmt.Printf("Adicionado -> %d. Tamanho: %d\n", value, l.Size)

}
func (l *LinkedList) AddAtEnd(value int) {
	newNode := &Node{Value: value}

	if l.IsEmpty() {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode

		l.Tail = newNode
	}
	l.Size++
	fmt.Printf("Adicionado -> %d. Tamanho: %d\n", value, l.Size)

}

func (l *LinkedList) RemoveAtStart() (int, error) {
	if l.IsEmpty() {
		return 0, fmt.Errorf("erro: a lista está vazia")
	}

	removedValue := l.Head.Value

	l.Head = l.Head.Next

	if l.Head == nil {
		l.Tail = nil
	}

	l.Size--

	fmt.Printf("Removido <- %d. Tamanho: %d\n", removedValue, l.Size)
	return removedValue, nil
}

func (l *LinkedList) RemoveAtEnd() (int, error) {
	if l.IsEmpty() {
		return 0, fmt.Errorf("erro: a lista está vazia")
	}

	removedValue := l.Tail.Value

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		current := l.Head

		for current.Next != l.Tail {
			current = current.Next
		}

		l.Tail = current
		l.Tail.Next = nil
	}

	l.Size--

	fmt.Printf("Removido <- %d. Tamanho: %d\n", removedValue, l.Size)

	return removedValue, nil
}

func (l *LinkedList) RemoveValue(value int) bool {
	if l.IsEmpty() {
		return false
	}

	if l.Head.Value == value {
		l.RemoveAtStart()
		return true
	}

	previous := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Value == value {
			previous.Next = current.Next

			if current == l.Tail {
				l.Tail = previous
			}
			l.Size--
			return true
		}
		previous = current
		current = current.Next

	}
	return false
}

func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedList) PrintList() {
	if l.IsEmpty() {
		fmt.Println("A lista está vazia")
		return
	}

	current := l.Head
	fmt.Print("Lista: ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Lista Ligada (como uma Fila Eficiente) ---")
	lista := LinkedList{}

	// Usando a lista como uma Fila
	lista.AddAtEnd(10)
	lista.AddAtEnd(20)
	lista.AddAtEnd(30)

	fmt.Println("-----------------------------------------")
	lista.PrintList()
	fmt.Println("-----------------------------------------")

	lista.RemoveAtStart()
	lista.RemoveAtStart()

	fmt.Println("-----------------------------------------")
	lista.PrintList()
	fmt.Println("-----------------------------------------")

	lista.AddAtEnd(40)
	lista.PrintList()
}
