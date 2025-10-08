package main

import (
	"container/list"
	"fmt"
)

const FixedArraySize = 16

type KeyValuePair struct {
	Key   string
	Value int
}

type HashTable struct {
	Buckets [FixedArraySize]*list.List
}

func NewHashTable() *HashTable {
	table := &HashTable{}

	for i := 0; i < FixedArraySize; i++ {
		table.Buckets[i] = list.New()
	}
	return table
}

func (t *HashTable) hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % FixedArraySize
}

func (t *HashTable) Insert(key string, value int) {
	index := t.hash(key)
	bucket := t.Buckets[index]

	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key {
			fmt.Printf("Atualizando chave '%s' no balde %d\n", key, index)
			pair.Value = value
			return
		}
	}

	fmt.Printf("Inserindo chave '%s' no balde %d\n", key, index)

	bucket.PushBack(&KeyValuePair{Key: key, Value: value})

}

func (t *HashTable) Remover(key string) bool {
	index := t.hash(key)
	bucket := t.Buckets[index]

	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key {
			bucket.Remove(e)
			fmt.Printf("Removendo chave '%s' do balde %d\n", key, index)

			return true
		}
	}
	return false
}

func (t *HashTable) Search(key string)(int, bool)  {
	index := t.hash(key)
	bucket := t.Buckets[index]

	for e := bucket.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*KeyValuePair)
		if pair.Key == key{
			return  pair.Value, true
		}		
	}
	return 0, false
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
