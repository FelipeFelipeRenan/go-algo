package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 16 // Altura máxima da nossa Skip List
const p = 0.5       // Probabilidade de um nó ser promovido para o próximo nível

// Node representa um nó na Skip List.
type Node struct {
	Value int
	// Forward é um slice de ponteiros para os próximos nós em cada nível.
	Forward []*Node
}

// SkipList é a estrutura principal.
type SkipList struct {
	Header *Node // Um nó "cabeçalho" sentinela para simplificar a lógica.
	Level  int   // O nível mais alto atualmente na lista.
}

// NewSkipList cria uma Skip List vazia.
func NewSkipList() *SkipList {
	// O cabeçalho aponta para nil em todos os níveis.
	header := &Node{Forward: make([]*Node, maxLevel)}
	return &SkipList{Header: header, Level: 1}
}

// randomLevel decide a "altura" de um novo nó.
func (sl *SkipList) randomLevel() int {
	level := 1
	// Continuamos "jogando a moeda" e aumentando o nível.
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// Inserir adiciona um novo valor à Skip List.
func (sl *SkipList) Insert(value int) {

	fmt.Printf("Inserindo '%d'...\n", value)
	// 'update' vai guardar os nós que precisam ser "atualizados" em cada nível.
	update := make([]*Node, maxLevel)
	currentNode := sl.Header

	// 1. Encontra a posição de inserção (e os nós a serem atualizados).
	// Começa do nível mais alto e desce.
	for i := sl.Level - 1; i >= 0; i-- {
		for currentNode.Forward[i] != nil && currentNode.Forward[i].Value < value {
			currentNode = currentNode.Forward[i] // Anda para frente no nível atual
		}
		update[i] = currentNode // Guarda o último nó antes da inserção no nível i
	}

	// 2. Decide a altura do novo nó.
	newLevel := sl.randomLevel()

	// 3. Se o novo nó for mais alto que a lista atual, atualiza o nível da lista.
	if newLevel > sl.Level {
		for i := sl.Level; i < newLevel; i++ {
			update[i] = sl.Header
		}
		sl.Level = newLevel
	}

	// 4. Cria e "costura" o novo nó na lista.
	newNode := &Node{Value: value, Forward: make([]*Node, newLevel)}

	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i] // O novo nó aponta para onde o anterior apontava
		update[i].Forward[i] = newNode            // O nó anterior agora aponta para o novo nó

	}
}

// Buscar procura por um valor na Skip List.
func (sl *SkipList) Search(value int) bool {
	currentNode := sl.Header

	// Começa do nível mais alto e desce, como na inserção.
	for i := sl.Level - 1; i >= 0; i-- {
		for currentNode.Forward[i] != nil && currentNode.Forward[i].Value < value {
			currentNode = currentNode.Forward[i]
		}
	}

	// Após o loop, estamos no nível 0, no nó anterior ao local onde o valor deveria estar.
	// Verificamos se o próximo nó de fato contém o valor.
	if currentNode.Forward[0] != nil && currentNode.Forward[0].Value == value {
		return true
	}
	return false
}

// Imprimir visualiza a Skip List.
func (sl *SkipList) PrintList() {
	fmt.Println("--- Visualização da Skip List ---")

	for i := sl.Level - 1; i >= 0; i-- {
		fmt.Printf("Nível %d: H -> ", i)
		node := sl.Header.Forward[i]
		for node != nil {
			fmt.Printf("%d ->  ", node.Value)
			node = node.Forward[i]
		}
		fmt.Println("nil")
	}
}

func main() {
	// Inicializa o gerador de números aleatórios.
	rand.Seed(time.Now().UnixNano())

	sl := NewSkipList()
	sl.Insert(3)
	sl.Insert(6)
	sl.Insert(7)
	sl.Insert(9)
	sl.Insert(12)
	sl.Insert(19)
	sl.Insert(17)
	sl.Insert(26)
	sl.Insert(21)
	sl.Insert(25)

	sl.PrintList()

	fmt.Println("\n--- Buscas ---")
	fmt.Printf("Valor 19 existe? %t\n", sl.Search(19))
	fmt.Printf("Valor 10 existe? %t\n", sl.Search(10))
}
