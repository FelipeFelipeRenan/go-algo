package main

import "fmt"


const AlphabetSize = 26 // Para este exemplo, usaremos apenas a-z minúsculo.

// TrieNode representa um nó na nossa Trie.
type TrieNode struct {
	// Filhos é um mapa onde a chave é o caractere e o valor é o nó filho.
	// Usar um mapa torna a implementação mais flexível do que um array.
	Childs map[rune]*TrieNode
	// FimDePalavra marca se o caminho até este nó forma uma palavra completa
	WordEnd bool
}

// Trie representa a arvore completa, só precisamos da raiz
type Trie struct {
	Root *TrieNode
}

// NewTrieNode cria um novo nó de Trie vazio.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Childs: make(map[rune]*TrieNode),
	}
}

// NewTrie cria uma Trie vazia com um nó raiz.
func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

// Insert adiciona uma palavra à Trie.
func (t *Trie) Insert(word string) {
	fmt.Printf("Inserindo '%s'...\n", word)

	currentNode := t.Root

	// Percorre cada caractere da palavra
	for _, char := range word {
		// Se o caractere não existe como filho do nó atual, crie-o.
		if _, exists := currentNode.Childs[char]; !exists {
			currentNode.Childs[char] = NewTrieNode()
		}
		// Desce para o próximo nó.
		currentNode = currentNode.Childs[char]
	}
	// Ao final da palavra, marca o último nó como um "fim de palavra".
	currentNode.WordEnd = true
}

// Search verifica se uma palavra exata existe na Trie.
func (t *Trie) Search(word string) bool {
	currentNode := t.Root

	for _, char := range word {
		// Se em qualquer ponto o caractere não existir como filho, a palavra não existe.
		if _, exists := currentNode.Childs[char]; !exists {
			return false
		}
		currentNode = currentNode.Childs[char]
	}

	// A palavra só existe se chegamos ao fim E o último nó está marcado como fim de palavra.
	// Isso diferencia "cart" de "carta".
	return currentNode.WordEnd
}

// StartsWith verifica se existe alguma palavra na Trie com um dado prefixo.
func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.Root

	for _, char := range prefix {
		if _, exists := currentNode.Childs[char]; !exists {
			return false
		}
		currentNode = currentNode.Childs[char]
	}
	// Se conseguimos percorrer todos os caracteres do prefixo, ele existe.
	return true
}

// --- Função Principal ---
func main() {
	minhaTrie := NewTrie()

	palavras := []string{"carro", "carta", "carteiro", "casa", "casaco"}
	for _, p := range palavras {
		minhaTrie.Insert(p)
	}

	fmt.Println("\n--- Buscando palavras exatas ---")
	fmt.Printf("A palavra 'carta' existe? %t\n", minhaTrie.Search("carta"))
	fmt.Printf("A palavra 'cart' existe? %t\n", minhaTrie.Search("cart")) // Falso, é só um prefixo
	fmt.Printf("A palavra 'casarão' existe? %t\n", minhaTrie.Search("casarão"))

	fmt.Println("\n--- Verificando prefixos ---")
	fmt.Printf("Existe alguma palavra com o prefixo 'cart'? %t\n", minhaTrie.StartsWith("cart"))
	fmt.Printf("Existe alguma palavra com o prefixo 'carr'? %t\n", minhaTrie.StartsWith("carr"))
	fmt.Printf("Existe alguma palavra com o prefixo 'gato'? %t\n", minhaTrie.StartsWith("gato"))
}
