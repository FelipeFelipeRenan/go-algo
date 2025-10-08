package main

import "fmt"

// MinHeap representa a nossa fila de prioridade.
type MinHeap struct {
	elements []int
}

// ---- Funções Auxiliares para Navegação ----
func (h *MinHeap) fatherIndex(i int) int     { return (i - 1) / 2 }
func (h *MinHeap) leftChildIndex(i int) int  { return 2*i + 1 }
func (h *MinHeap) rightChildIndex(i int) int { return 2*i + 2 }

func (h *MinHeap) switchElements(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// ---- Lógica Principal: Borbulhar para Cima e para Baixo ----

// heapifyUp (borbulhar para cima) é usado após a inserção.
// Garante que o novo elemento suba até sua posição correta.
func (h *MinHeap) heapfyUp() {
	index := len(h.elements) - 1 // Começa com o último elemento
	for index > 0 && h.elements[h.fatherIndex(index)] > h.elements[index] {
		father := h.fatherIndex(index)
		fmt.Printf("   ...heapifyUp: trocando %d (índice %d) com pai %d (índice %d)\n", h.elements[index], index, h.elements[father], father)
		h.switchElements(father, index)
		index = father // Move para a posição do pai para continuar a verificação
	}
}

// heapifyDown (borbulhar para baixo) é usado após a extração.
// Garante que o elemento que substituiu a raiz desça até sua posição correta.
func (h *MinHeap) heapfyDown() {
	index := 0 // Começa com a raiz
	size := len(h.elements)

	// Continua enquanto houver pelo menos um filho à esquerda
	for h.leftChildIndex(index) < size {
		smallerChildIndex := h.leftChildIndex(index)
		rightIndex := h.rightChildIndex(index)

		// Verifica se o filho da direita existe e é menor que o da esquerda
		if rightIndex < size && h.elements[rightIndex] < h.elements[smallerChildIndex] {
			smallerChildIndex = rightIndex
		}

		// Se o pai já for menor que o menor dos filhos, a propriedade está ok.
		if h.elements[index] <= h.elements[smallerChildIndex] {
			break
		}
		fmt.Printf("   ...heapifyDown: trocando %d (índice %d) com filho %d (índice %d)\n", h.elements[index], index, h.elements[smallerChildIndex], smallerChildIndex)
		h.switchElements(index, smallerChildIndex)
		index = smallerChildIndex

	}
}

// Inserir adiciona um elemento à heap. Complexidade: O(log n)
func (h *MinHeap) Insert(value int) {
	fmt.Printf("Inserindo %d ...\n", value)
	h.elements = append(h.elements, value)
	h.heapfyUp()
	fmt.Printf("  ...Heap atual: %v\n", h.elements)
}

// ExtrairMin remove e retorna o menor elemento. Complexidade: O(log n)
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap vazio")
	}

	min := h.elements[0]
	last := len(h.elements) - 1

	fmt.Printf("Extraindo o minimo (%d) ...\n", min)

	// Move o último elemento para a raiz
	h.elements[0] = h.elements[last]
	// E remove o último elemento (encurtando o slice)
	h.elements = h.elements[:last]
	fmt.Printf("   Heap antes do heapifyDown: %v\n", h.elements)
	h.heapfyDown()
	fmt.Printf("   Heap após o heapifyDown: %v\n", h.elements)
	return min, nil
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Min-Heap ---")
	minhaHeap := &MinHeap{}

	minhaHeap.Insert(10)
	minhaHeap.Insert(4)
	minhaHeap.Insert(15)
	minhaHeap.Insert(20)
	minhaHeap.Insert(1) // Este vai "borbulhar" até o topo

	fmt.Println("\n--- Extraindo elementos em ordem de prioridade ---")
	for i := 0; i < 5; i++ {
		min, err := minhaHeap.ExtractMin()
		if err == nil {
			fmt.Printf("--> Mínimo extraído: %d\n\n", min)
		}
	}

	// Heapsort
	fmt.Println("--- Demonstração do Heapsort ---")
	numeros := []int{4, 10, 3, 5, 1, 12}
	fmt.Printf("Slice original: %v\n\n", numeros)

	Heapsort(numeros)

	fmt.Printf("\nSlice ordenado: %v\n", numeros)

}
