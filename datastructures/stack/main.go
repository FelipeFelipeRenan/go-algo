package main

import "fmt"

type Stack struct {
	elementos []int
}

func (s *Stack) Push(elemento int) {
	s.elementos = append(s.elementos, elemento)
	fmt.Printf("Pushed -> %d. Pilha atual: %v\n", elemento, s.elementos)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("erro: a pilha está vazia")
	}

	indiceDoTopo := len(s.elementos) - 1

	elementoDoTopo := s.elementos[indiceDoTopo]

	s.elementos = s.elementos[:indiceDoTopo]

	fmt.Printf("Pop <- %d. Pilha atual: %v\n", elementoDoTopo, s.elementos)
	return elementoDoTopo, nil

}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("erro: a pilha está vazia")
	}
	return s.elementos[len(s.elementos)-1], nil
}

func (s *Stack) Size() int {
	return len(s.elementos)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elementos) == 0
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Pilha ---")
	minhaPilha := Stack{} // Cria uma nova instância da nossa pilha

	fmt.Printf("A pilha está vazia? %t\n", minhaPilha.IsEmpty())
	fmt.Println("-----------------------------------------")

	// Empilhando alguns números
	minhaPilha.Push(10)
	minhaPilha.Push(20)
	minhaPilha.Push(30)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Tamanho da pilha: %d\n", minhaPilha.Size())
	fmt.Printf("A pilha está vazia? %t\n", minhaPilha.IsEmpty())

	elementoTopo, _ := minhaPilha.Peek()
	fmt.Printf("Olhando o topo (Peek): %d\n", elementoTopo)
	fmt.Println("-----------------------------------------")

	// Desempilhando
	minhaPilha.Pop()
	minhaPilha.Pop()
	minhaPilha.Pop()

	fmt.Println("-----------------------------------------")
	fmt.Printf("Tamanho final da pilha: %d\n", minhaPilha.Size())

	// Tentando dar Pop em uma pilha vazia
	_, err := minhaPilha.Pop()
	if err != nil {
		fmt.Printf("Erro esperado: %s\n", err)
	}
}