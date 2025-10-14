package main

import "fmt"

// SegmentTree é a nossa estrutura.
type SegmentTree struct {
	original []int // O array original, para referência
	tree     []int // A árvore em si, armazenada em um slice
}

// NewSegmentTree cria e constrói a árvore a partir de um slice de dados.
func NewSegmentTree(data []int) *SegmentTree {
	st := &SegmentTree{
		original: data,
		// A árvore precisa de até 4x o espaço do array original para caber.
		tree: make([]int, 4*len(data)),
	}
	// Inicia o processo de construção recursiva.
	st.build(0, 0, len(data)-1)
	return st
}

// build é a função recursiva que constrói a árvore.
// treeIdx: índice do nó atual na nossa 'tree'.
// start, end: o intervalo do array original que este nó representa.
func (st *SegmentTree) build(treeIdx, start, end int) {
	// Caso base: chegamos a uma folha, que representa um único elemento.
	if start == end {
		st.tree[treeIdx] = st.original[start]
		return
	}
	// Passo recursivo:
	mid := (start + end) / 2
	leftTreeIdx := 2*treeIdx + 1
	rightTreeIdx := 2*treeIdx + 2

	// Constrói a sub-árvore esquerda e a direita.
	st.build(leftTreeIdx, start, mid)
	st.build(rightTreeIdx, mid+1, end)

	// O valor do nó atual é a soma dos seus dois filhos.
	st.tree[treeIdx] = st.tree[leftTreeIdx] + st.tree[rightTreeIdx]
}

// Query é a função pública para fazer uma consulta de intervalo.
func (st *SegmentTree) Query(left, right int) int {
	return st.query(0, 0, len(st.original)-1, left, right)
}

// query é a função recursiva que encontra a soma do intervalo [queryL..queryR].
func (st *SegmentTree) query(treeIdx, start, end, queryL, queryR int) int {
	// Caso 1: O intervalo do nó está totalmente fora do intervalo da consulta.
	// Não contribui com nada para a soma.
	if end < queryL || start > queryR {
		return 0
	}
	// Caso 2: O intervalo do nó está totalmente dentro do intervalo da consulta.
	// Retornamos o valor pré-calculado deste nó.
	if queryL <= start && end <= queryR {
		return st.tree[treeIdx]
	}
	// Caso 3: O intervalo do nó se sobrepõe parcialmente ao da consulta.
	// Precisamos perguntar aos filhos e somar os resultados.
	mid := (start + end) / 2
	leftSum := st.query(2*treeIdx+1, start, mid, queryL, queryR)
	rightSum := st.query(2*treeIdx+2, mid+1, end, queryL, queryR)
	return leftSum + rightSum
}

// Update é a função pública para atualizar um valor no array original.
func (st *SegmentTree) Update(index, newValue int) {
	st.original[index] = newValue
	st.update(0, 0, len(st.original)-1, index, newValue)
}

// update é a função recursiva que atualiza a árvore.
func (st *SegmentTree) update(treeIdx, start, end, updateIdx, value int) {
	// Caso base: encontramos a folha correspondente ao índice.
	if start == end {
		st.tree[treeIdx] = value
		return
	}

	// Passo recursivo: determina se vai para a esquerda ou direita.
	mid := (start + end) / 2

	if updateIdx <= mid {
		st.update(2*treeIdx+1, start, mid, updateIdx, value)
	} else {
		st.update(2*treeIdx+2, mid+1, end, updateIdx, value)
	}
	// Após a atualização do filho, recalcula o valor do nó atual.
	// É assim que a mudança "borbulha" para cima.
	st.tree[treeIdx] = st.tree[2*treeIdx+1] + st.tree[2*treeIdx+2]
}

func main() {
	dados := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("Array Original: %v\n", dados)

	st := NewSegmentTree(dados)
	fmt.Printf("Árvore de Segmentos (slice interno): %v\n", st.tree)

	fmt.Println("\n--- Consultas ---")
	// Soma dos elementos do índice 1 ao 3 (3 + 5 + 7)
	soma := st.Query(1, 3)
	fmt.Printf("A soma do intervalo [1..3] é: %d\n", soma) // Esperado: 15

	// Soma de todo o array
	somaTotal := st.Query(0, 5)
	fmt.Printf("A soma de todo o array [0..5] é: %d\n", somaTotal) // Esperado: 36

	fmt.Println("\n--- Atualização ---")
	// Atualizando o elemento no índice 2 de 5 para 6
	fmt.Println("Atualizando o índice 2 de 5 para 6...")
	st.Update(2, 6)
	fmt.Printf("Novo Array Original: %v\n", st.original)

	// Fazendo a mesma consulta de antes
	soma = st.Query(1, 3)
	fmt.Printf("A NOVA soma do intervalo [1..3] é: %d\n", soma) // Esperado: 3 + 6 + 7 = 16
}
