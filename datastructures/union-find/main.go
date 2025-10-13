package main

import "fmt"

// Disjoint Set Union
type DSU struct {
	// O mapa 'parent' armazena o pai de cada elemento.
	parent map[int]int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make(map[int]int),
	}
	// NewDSU cria e inicializa a estrutura para um conjunto de elementos.
	// Neste caso, de 0 a n-1.
	for i := 0; i < n; i++ {
		// No início, cada elemento é seu próprio pai (está em seu próprio grupo).

		dsu.parent[i] = i
	}
	return dsu
}

// Find encontra o representante (a raiz) do conjunto ao qual 'i' pertence.
// Esta implementação inclui a otimização de "Compressão de Caminho".
func (d *DSU) Find(i int) int {
	// Se 'i' é seu próprio pai, ele é a raiz.
	if d.parent[i] == i {
		return i
	}

	// Se não, sobe recursivamente para encontrar a raiz...
	// No caminho de volta da recursão, fazemos com que 'i' aponte
	// DIRETAMENTE para a raiz que encontramos.
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

func (d *DSU) Union(i, j int) {
	// Encontra os representantes de ambos os elementos.
	rootI := d.Find(i)
	rootJ := d.Find(j)

	// Se eles não estiverem no mesmo conjunto, une os dois.
	if rootI != rootJ {

		// Simplesmente faz com que o representante de um grupo aponte para o outro.
		fmt.Printf("Unindo grupos de %d e %d. Novo representante: %d\n", i, j, rootJ)
		d.parent[rootI] = rootJ
	} else {
		fmt.Printf("%d e %d já estão no mesmo grupo.\n", i, j)

	}
}

// --- Função Principal ---
func main() {
	fmt.Println("--- Iniciando a demonstração da Union-Find (DSU) ---")
	// Criamos um conjunto com 5 pessoas (0 a 4)
	dsu := NewDSU(5)

	fmt.Printf("0 e 4 estão no mesmo grupo? %t\n", dsu.Find(0) == dsu.Find(4)) // false

	dsu.Union(0, 1) // 0 e 1 se tornam amigos
	dsu.Union(2, 3) // 2 e 3 se tornam amigos

	fmt.Printf("0 e 1 estão no mesmo grupo? %t\n", dsu.Find(0) == dsu.Find(1)) // true
	fmt.Printf("0 e 2 estão no mesmo grupo? %t\n", dsu.Find(0) == dsu.Find(2)) // false

	fmt.Println("\n--- Unindo dois grupos existentes ---")
	dsu.Union(1, 3) // O amigo de 0 (1) se torna amigo do amigo de 2 (3)

	// Agora, por transitividade, 0 e 2 devem estar no mesmo grupo!
	fmt.Printf("AGORA, 0 e 2 estão no mesmo grupo? %t\n", dsu.Find(0) == dsu.Find(2)) // true
	fmt.Printf("E 0 e 3? %t\n", dsu.Find(0) == dsu.Find(3))                           // true

	// Vamos tentar unir dois que já estão no mesmo grupo
	dsu.Union(0, 3)

	
}
