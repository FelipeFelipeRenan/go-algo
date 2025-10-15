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

// Remover inicia o processo de remoção de uma palavra.
func (t *Trie) Remove(word string) {
	fmt.Printf("\nTentando remover '%s'", word)
	t.removeRecursive(t.Root, word, 0)
}

// removerRecursivo faz a remoção de forma recursiva.
// Retorna 'true' se o nó pai deve remover o ponteiro para o nó atual.
func (t *Trie) removeRecursive(currentNode *TrieNode, word string, depth int) bool {
	// Caso base: se o nó é nulo, não há nada a fazer.
	if currentNode == nil {
		return false
	}
	// Caso recursivo: chegamos ao final da palavra.
	if depth == len(word) {
		// Se o nó marca um fim de palavra, desmarcamos.
		if currentNode.WordEnd {
			currentNode.WordEnd = false
			// Se este nó não tiver mais filhos, ele se tornou inútil
			// e pode ser removido pelo seu pai.
			return len(currentNode.Childs) == 0
		}
		// Se não era um fim de palavra, não há nada a remover.
		return false
	}

	// Passo recursivo: desce na árvore.
	char := rune(word[depth])
	if shouldRemoveChild := t.removeRecursive(currentNode.Childs[char], word, depth+1); shouldRemoveChild {
		// Se a chamada recursiva disse para remover o filho, nós o removemos.
		delete(currentNode.Childs, char)
		// E verificamos se o nó atual também se tornou inútil.
		return !currentNode.WordEnd && len(currentNode.Childs) == 0
	}
	return false
}

// EncontrarPorPrefixo retorna um slice com todas as palavras que começam com o prefixo.
func (t *Trie) FindByPrefix(prefix string) []string {

	currentNode := t.Root

	// 1. Navega até o final do prefixo
	for _, char := range prefix {
		if _, exists := currentNode.Childs[char]; !exists {
			// Se o prefixo não existe, não há palavras.
			return []string{}
		}
		currentNode = currentNode.Childs[char]
	}
	// 2. Coleta todas as palavras daquele ponto em diante
	var wordsFound []string
	t.collectWords(currentNode, prefix, &wordsFound)
	return wordsFound
}

// coletarPalavras é uma função DFS recursiva para encontrar as palavras.
func (t *Trie) collectWords(currentNode *TrieNode, currentPrefix string, words *[]string) {
	// Se o nó atual é um fim de palavra, adicionamos à nossa lista.
	if currentNode.WordEnd {
		*words = append(*words, currentPrefix)
	}
	// Para cada filho, continua a busca recursivamente.
	for char, childNode := range currentNode.Childs {
		t.collectWords(childNode, currentPrefix+string(char), words)
	}
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

	fmt.Println("\n--- Testando o Auto-Completar ---")
	prefixo := "cart"
	sugestoes := minhaTrie.FindByPrefix(prefixo)
	fmt.Printf("Palavras que começam com '%s': %v\n", prefixo, sugestoes)

	prefixo2 := "casa"
	sugestoes2 := minhaTrie.FindByPrefix(prefixo2)
	fmt.Printf("Palavras que começam com '%s': %v\n", prefixo2, sugestoes2)

	fmt.Println("\n--- Testando a Remoção ---")
	fmt.Printf("A palavra 'carteiro' existe? %t\n", minhaTrie.Search("carteiro"))
	minhaTrie.Remove("carteiro")
	fmt.Printf("A palavra 'carteiro' existe agora? %t\n", minhaTrie.Search("carteiro"))
	fmt.Printf("A palavra 'carta' ainda existe? %t\n", minhaTrie.Search("carta")) // Deve ser true

	sugestoesAposRemocao := minhaTrie.FindByPrefix(prefixo)
	fmt.Printf("Sugestões para '%s' após remoção: %v\n", prefixo, sugestoesAposRemocao)

}
