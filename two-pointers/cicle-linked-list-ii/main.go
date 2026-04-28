package main

import "fmt"

func main() {
	// Exemplo de criação de ciclo manual para teste
	n1 := &Node{Value: 3}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 0}
	n4 := &Node{Value: -4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // Ciclo começa no nó com valor 2

	result := DetectCycle(n1)
	if result != nil {
		fmt.Printf("Ciclo encontrado no nó com valor: %d\n", result.Value)
	} else {
		fmt.Println("Nenhum ciclo encontrado.")
	}
}
