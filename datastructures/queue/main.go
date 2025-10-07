package main

import "fmt"

type Queue struct {
	elementos []int
}

func (q *Queue) Enqueue(elemento int) {
	q.elementos = append(q.elementos, elemento)
	fmt.Printf("Enqueue -> %d. Fila atual: %v\n", elemento, q.elementos)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("erro: a fila está vazia")
	}

	elementoDoInicio := q.elementos[0]

	q.elementos = q.elementos[1:]
	fmt.Printf("Dequeue <- %d. Fila atual: %v\n", elementoDoInicio, q.elementos)
	return elementoDoInicio, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("erro: a fila está vazia")
	}
	return q.elementos[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elementos) == 0
}

func (q *Queue) Size() int {
	return len(q.elementos)
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Fila ---")
	minhaFila := Queue{}

	fmt.Printf("A fila está vazia? %t\n", minhaFila.IsEmpty())
	fmt.Println("-----------------------------------------")

	// Enfileirando alguns números
	minhaFila.Enqueue(100)
	minhaFila.Enqueue(200)
	minhaFila.Enqueue(300)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Tamanho da fila: %d\n", minhaFila.Size())

	elementoInicio, _ := minhaFila.Peek()
	fmt.Printf("Olhando o início (Peek): %d\n", elementoInicio)
	fmt.Println("-----------------------------------------")

	// Desenfileirando
	minhaFila.Dequeue()
	minhaFila.Dequeue()

	fmt.Println("-----------------------------------------")
	fmt.Printf("Fila final: %v\n", minhaFila.elementos)
	fmt.Printf("Tamanho final da fila: %d\n", minhaFila.Size())

	minhaFila.Dequeue()
	_, err := minhaFila.Dequeue()
	if err != nil {
		fmt.Printf("Erro esperado: %s\n", err)
	}
}
