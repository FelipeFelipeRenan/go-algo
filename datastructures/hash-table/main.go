package main

import (
	"container/list"
	"fmt"
)

// TamanhoFixoArray define o número de baldes em nossa tabela.
// Em uma implementação real, isso seria dinâmico.
const FixedArraySize = 16

// KeyVAluePair guarda a informação que será armazenada na lista ligada.
type KeyValuePair struct {
	Key   string
	Value int
}

// HashTable é a nossa estrutura principal
type HashTable struct {
	Buckets [FixedArraySize]*list.List
}

// função para criar nova tabela hash
func NewHashTable() *HashTable {
	table := &HashTable{}

	for i := 0; i < FixedArraySize; i++ {
		// Inicializa a lista em cada balde
		table.Buckets[i] = list.New()
	}
	return table
}

// hash é nossa função de hash simples para strings.
// Ela converte a chave em um índice para o nosso array de baldes.
func (t *HashTable) hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	// O operador '%' garante que o índice esteja dentro dos limites do nosso array.
	return sum % FixedArraySize
}

func (t *HashTable) Insert(key string, value int) {
	index := t.hash(key)
	bucket := t.Buckets[index]

	// Passo 1: Verifica se a chave já existe. Se sim, atualiza o valor.
	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key {
			fmt.Printf("Atualizando chave '%s' no balde %d\n", key, index)
			pair.Value = value
			return
		}
	}

	// Passo 2: Se a chave não existe, adiciona no final da lista ligada.
	fmt.Printf("Inserindo chave '%s' no balde %d\n", key, index)

	bucket.PushBack(&KeyValuePair{Key: key, Value: value})

}

// Remover deleta um par chave-valor.
func (t *HashTable) Remover(key string) bool {
	index := t.hash(key)
	bucket := t.Buckets[index]

	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key {
			bucket.Remove(e) // Remove o elemento da lista ligada
			fmt.Printf("Removendo chave '%s' do balde %d\n", key, index)

			return true
		}
	}
	return false
}

// Buscar encontra um valor associado a uma chave.
func (t *HashTable) Search(key string) (int, bool) {
	index := t.hash(key)
	bucket := t.Buckets[index]

	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key {
			return pair.Value, true // Encontrado
		}
	}
	return 0, false // Não encontrado 
}

func main() {
	fmt.Println("--- Iniciando a demonstração da Tabela Hash ---")
	minhaTabela := NewHashTable()

	// Inserindo valores
	minhaTabela.Insert("uva", 10)
	minhaTabela.Insert("maçã", 20)
	minhaTabela.Insert("banana", 30)
	// "abacaxi" e "uva" podem colidir dependendo da função de hash e tamanho
	minhaTabela.Insert("abacaxi", 40)

	fmt.Println("\n--- Realizando Buscas ---")
	valor, ok := minhaTabela.Search("maçã")
	if ok {
		fmt.Printf("Valor de 'maçã': %d\n", valor)
	}

	valor, ok = minhaTabela.Search("banana")
	if ok {
		fmt.Printf("Valor de 'banana': %d\n", valor)
	}

	valor, ok = minhaTabela.Search("laranja")
	if !ok {
		fmt.Println("Chave 'laranja' não encontrada (esperado).")
	}

	fmt.Println("\n--- Atualizando e Removendo ---")
	minhaTabela.Insert("maçã", 25) // Atualiza o valor
	valor, _ = minhaTabela.Search("maçã")
	fmt.Printf("Novo valor de 'maçã': %d\n", valor)

	minhaTabela.Remover("uva")
	_, ok = minhaTabela.Search("uva")
	if !ok {
		fmt.Println("Chave 'uva' não encontrada após remoção (esperado).")
	}
}
