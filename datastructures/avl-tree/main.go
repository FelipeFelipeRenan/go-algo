package main

import "fmt"

// AVLNode é o nosso nó, agora com o campo Altura.
type AVLNode struct {
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// AVLTree é a estrutura principal.
type AVLTree struct {
	Root *AVLNode
}

// --- Funções Auxiliares ---

// altura retorna a altura de um nó (trata o caso de nó nulo).
func height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// max retorna o maior de dois inteiros.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// balanco calcula o fator de balanço de um nó.
func balance(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// Rotações

// rightRotation executa uma rotação simples à direita.
func (t *AVLTree) rightRotation(y *AVLNode) *AVLNode {
	fmt.Println("... Executando rotação à direita em ", y.Value)
	x := y.Left
	T2 := x.Right

	// Realiza a rotação
	x.Right = y
	y.Left = T2

	// Atualiza as alturas
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x // x é a nova raiz da sub-árvore
}

// leftRotation executa uma rotação simples à esquerda.
func (t *AVLTree) leftRotation(x *AVLNode) *AVLNode {
	fmt.Println("... Executando rotação à direita em ", x.Value)
	y := x.Right
	T2 := y.Left

	// Realiza a rotação
	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y // y é a nova raiz da sub-árvore
}

// --- Inserção ---

// Inserir é o método público que inicia a inserção.
func (t *AVLTree) Insert(value int) {
	fmt.Printf("Inserindo %d... \n", value)
	t.Root = t.insert(t.Root, value)
}

// inserir é a função recursiva que insere e balanceia a árvore.
func (t *AVLTree) insert(node *AVLNode, value int) *AVLNode {
	// 1. Inserção normal de BST
	if node == nil {
		return &AVLNode{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = t.insert(node.Left, value)
	} else if value > node.Value {
		node.Right = t.insert(node.Right, value)
	} else {
		return node // Valores duplicados não são permitidos
	}

	// 2. Atualiza a altura do nó atual (ancestral)
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// 3. Calcula o fator de balanço e verifica se precisa de rotação
	balanceFactor := balance(node)

	// Caso Esquerda-Esquerda (EE)
	if balanceFactor > 1 && value < node.Left.Value {
		return t.rightRotation(node)
	}
	// Caso Direita-Direita (DD)
	if balanceFactor < -1 && value > node.Right.Value {
		return t.leftRotation(node)
	}
	// Caso Esquerda-Direita (ED)
	if balanceFactor > 1 && value > node.Left.Value {
		node.Left = t.leftRotation(node.Left)
		return t.rightRotation(node)
	}
	// Caso Direita-Esquerda (DE)
	if balanceFactor < -1 && value < node.Right.Value {
		node.Right = t.rightRotation(node.Right)
		return t.leftRotation(node)
	}
	// 4. Se não precisou rotacionar, retorna o nó sem alterações.
	return node
}

// PercorrerEmOrdem para verificar se a árvore ainda é uma BST válida.
func (t *AVLTree) InOrderTraverse(node *AVLNode) {
	if node != nil {
		t.InOrderTraverse(node.Left)
		fmt.Printf("%d ", node.Value)
		t.InOrderTraverse(node.Right)
	}
}

func main() {
	avl := &AVLTree{}

	fmt.Println("--- Inserindo em ordem (o pior caso para uma BST normal) ---")
	avl.Insert(10)
	avl.Insert(20)

	// A inserção do 30 vai causar um desbalanceamento DD
	avl.Insert(30)

	fmt.Println("\n--- Inserindo mais valores ---")
	avl.Insert(40)

	// A inserção do 50 vai causar outro desbalanceamento DD
	avl.Insert(50)

	// A inserção do 25 vai causar um desbalanceamento DE
	avl.Insert(25)

	fmt.Printf("\n\nTravessia Em Ordem (prova de que a árvore está correta): ")
	avl.InOrderTraverse(avl.Root)
	fmt.Println()
}
