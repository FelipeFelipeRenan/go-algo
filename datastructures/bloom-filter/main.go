package main

import (
	"fmt"
	"hash/crc32"
	"hash/fnv"
)

type BloomFilter struct {
	bits []bool
	k    int
	m    int
}

func NewBloomFilter(m, k int) *BloomFilter {
	return &BloomFilter{
		bits: make([]bool, m),
		k:    k,
		m:    m,
	}
}


func (bf *BloomFilter) hash(item string) []int {

	hashes := make([]int, bf.k)

	h1 := fnv.New32a()
	h1.Write([]byte(item))
	hash1 := int(h1.Sum32())

	h2 := crc32.NewIEEE()
	h2.Write([]byte(item))
	hash2 := int(h2.Sum32())

	for i := 0; i < bf.k; i++ {

		combinedHash := hash1 + i*hash2

		hashes[i] = (combinedHash &  0x7fffffff) % bf.m
	}

	return hashes
}

func (bf *BloomFilter) Insert(item string) {
	fmt.Printf("Inserindo '%s'...\n", item)
	hashes := bf.hash(item)
	for _, h := range hashes {
		bf.bits[h] = true
	}
}

func (bf *BloomFilter) Search(item string) bool {
	hashes := bf.hash(item)
	for _, h := range hashes {

		if !bf.bits[h] {
			return false
		}
	}

	return true
}

// --- Função Principal ---
func main() {
	// Criamos um filtro com um array de 100 bits e 3 funções de hash.
	filtro := NewBloomFilter(100, 3)

	filtro.Insert("goiaba")
	filtro.Insert("laranja")

	fmt.Println("\n--- Verificando a existência ---")
	fmt.Printf("A fruta 'goiaba' existe? %t\n", filtro.Search("goiaba"))   // True (Certeza)
	fmt.Printf("A fruta 'laranja' existe? %t\n", filtro.Search("laranja")) // True (Certeza)
	fmt.Printf("A fruta 'morango' existe? %t\n", filtro.Search("morango")) // False (Certeza)

	// 'abacaxi' não foi inserido, mas pode dar um falso positivo
	// (a chance é baixa com estes parâmetros).
	fmt.Printf("A fruta 'abacaxi' existe? %t\n", filtro.Search("abacaxi")) // Provavelmente False
}
