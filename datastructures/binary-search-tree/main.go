package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = bst.InsertNode(bst.Root, value)
}

func (bst *BinarySearchTree) InsertNode(currentNode *Node, value int) *Node {

	if currentNode == nil {
		fmt.Printf("Inserindo %d\n", value)
		return &Node{Value: value}
	}

	if value < currentNode.Value {
		fmt.Printf("...Valor %d < %d, indo para a esquerda\n", value, currentNode.Value)
		currentNode.Left = bst.InsertNode(currentNode.Left, value)
	} else if value > currentNode.Value {
		fmt.Printf("...Valor %d < %d, indo para a esquerda\n", value, currentNode.Value)
		currentNode.Right = bst.InsertNode(currentNode.Right, value)
	}

	return currentNode
}

func (bst *BinarySearchTree) Search(value int) bool {
	return bst.SearchNode(bst.Root, value)
}

func (bst *BinarySearchTree) SearchNode(currentNode *Node, value int) bool {

	if currentNode == nil {
		return false
	}

	if value == currentNode.Value {
		return true
	}

	if value < currentNode.Value {
		return bst.SearchNode(currentNode.Left, value)
	}

	return bst.SearchNode(currentNode.Right, value)
}

// 1. Travessia Em Ordem (In-Order)
// Lógica: Esquerda -> Raiz -> Direita
// Resultado: Imprime os valores em ORDEM CRESCENTE.
func (bst *BinarySearchTree) InOrderTraverse() {
	fmt.Print("Em Ordem (In-Order):  ")
	bst.inOrderTraverse(bst.Root)
	fmt.Println()
}

func (bst *BinarySearchTree) inOrderTraverse(currentNode *Node) {
	if currentNode != nil {
		bst.inOrderTraverse(currentNode.Left)
		fmt.Printf("%d ", currentNode.Value)
		bst.inOrderTraverse(currentNode.Right)
	}
}

// 2. Travessia Pré-Ordem (Pre-Order)
// Lógica: Raiz -> Esquerda -> Direita
// Resultado: Útil para copiar uma árvore, pois recria a mesma estrutura.
func (bst *BinarySearchTree) PreOrderTraverse() {
	fmt.Print("Pré-Ordem (Pre-Order):  ")
	bst.preOrderTraverse(bst.Root)
	fmt.Println()
}

func (bst *BinarySearchTree) preOrderTraverse(currentNode *Node) {
	if currentNode != nil {
		fmt.Printf("%d ", currentNode.Value)
		bst.preOrderTraverse(currentNode.Left)
		bst.preOrderTraverse(currentNode.Right)
	}
}

// 3. Travessia Pós-Ordem (Post-Order)
// Lógica: Esquerda -> Direita -> Raiz
// Resultado: Útil para deletar uma árvore, pois deleta os filhos antes do pai.
func (bst *BinarySearchTree) PostOrderTraverse() {
	fmt.Print("Pós-Ordem (Post-Order):  ")
	bst.postOrderTraverse(bst.Root)
	fmt.Print()
}

func (bst *BinarySearchTree) postOrderTraverse(currentNode *Node) {
	if currentNode != nil {
		bst.postOrderTraverse(currentNode.Left)
		bst.postOrderTraverse(currentNode.Right)
		fmt.Printf("%d ", currentNode.Value)
	}
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Árvore de Busca Binária ---")
	arvore := BinarySearchTree{}
	valores := []int{8, 3, 10, 1, 6, 14, 4}
	for _, v := range valores {
		arvore.Insert(v)
		fmt.Println("--------------")
	}

	/* A árvore resultante se parece com isto:
	        8
	      /   \
	     3     10
	    / \      \
	   1   6      14
	      /
	     4
	*/

	fmt.Println("\n--- Realizando Buscas ---")
	fmt.Printf("O valor 6 existe na árvore? %t\n", arvore.Search(6))
	fmt.Printf("O valor 1 existe na árvore? %t\n", arvore.Search(1))
	fmt.Printf("O valor 99 existe na árvore? %t\n", arvore.Search(99))
	fmt.Printf("O valor 14 existe na árvore? %t\n", arvore.Search(14))

	fmt.Println("\n--- Demonstração das Travessias ---")
	
	arvore.InOrderTraverse()
	arvore.PreOrderTraverse()
	arvore.PostOrderTraverse()
}
