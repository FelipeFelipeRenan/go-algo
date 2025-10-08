package main

import "fmt"

// Heapsort é a função principal que ordena o slice.
func Heapsort(slice []int) {

	size := len(slice)

	// --- FASE 1: Construir a Max-Heap ---
	// Começamos do último nó pai e vamos até a raiz, aplicando heapifyDown
	// para garantir que cada sub-árvore seja uma Max-Heap.
	fmt.Println("--- Fase 1: Construindo a Max-Heap ---")
	for i := size/2 - 1; i >= 0; i-- {
		heapfyDownMax(slice, size, i)
	}
	fmt.Printf("Slice após a construção da Max-Heap: %v\n", slice)

	// --- FASE 2: Extrair elementos da heap para ordenar ---
	fmt.Println("--- Fase 2: Ordenando o Slice ---")
	for i := size - 1; i > 0; i-- {
		// Move a raiz atual (o maior elemento) para o fim.
		fmt.Printf("Trocando raiz %d com o fim da heap %d\n", slice[0], slice[i])
		slice[0], slice[i] = slice[i], slice[0]

		// Chama heapifyDownMax na heap reduzida para consertá-la.
		// A heap agora tem tamanho 'i'.
		heapfyDownMax(slice, i, 0)
		fmt.Printf("   Slice atual: %v\n", slice)
	}
}

// heapifyDownMax garante que a sub-árvore com raiz em 'raizIdx'
// obedeça à propriedade da Max-Heap. 'tamanhoHeap' define o limite da heap.
func heapfyDownMax(slice []int, size, rootIndex int) {

	largestIndex := rootIndex
	leftIndex := 2*rootIndex + 1
	rightIndex := 2*rootIndex + 2

	// Verifica se o filho da esquerda existe e é maior que a raiz
	if leftIndex < size && slice[leftIndex] > slice[largestIndex] {
		largestIndex = leftIndex
	}

	// Verifica se o filho da direita existe e é maior que o 'maior' encontrado até agora
	if rightIndex < size && slice[rightIndex] > slice[largestIndex] {
		largestIndex = rightIndex
	}
	// Se o maior elemento não for a raiz, troca e continua o processo recursivamente
	if largestIndex != rootIndex {
		slice[rootIndex], slice[largestIndex] = slice[largestIndex], slice[rootIndex]
		// Continua a "afundar" o elemento que estava na raiz original
		heapfyDownMax(slice, size, largestIndex)
	}
}
