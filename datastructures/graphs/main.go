package main

import (
	"container/heap"
	"fmt"
	"math"
)

// --- Estruturas para a Fila de Prioridade (Min-Heap) ---

// Item representa um item na nossa fila de prioridade.
// Precisamos saber o vértice, a prioridade (peso) e o índice na heap.
type Item struct {
	Vertex int
	Weight int // Usaremos o peso como prioridade
	index  int // Posição do item na heap
}

// PriorityQueue implementa a interface heap.Interface.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Weight < pq[j].Weight }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // evita vazamento de memória
	item.index = -1 // para segurança
	*pq = old[0 : n-1]
	return item
}

// Struct que representa uma aresta ponderada
type Edge struct {
	Destination int
	Weight      int
}

// Grafo representa a estrutura de um grafo usando uma lista de adjacência.
type Graph struct {
	// O mapa onde a chave é um vértice (int) e o valor é um slice
	// de todos os vértices vizinhos a ele.
	adjacencies map[int][]Edge
}

// NovoGrafo cria e inicializa um grafo vazio.
func NewGraph() *Graph {
	return &Graph{
		adjacencies: make(map[int][]Edge),
	}
}

// AdicionarVertice garante que um vértice exista no mapa, mesmo que sem arestas.
func (g *Graph) AddVertex(vertex int) {
	// Se o vértice ainda não estiver no mapa, adiciona com uma lista vazia.
	if _, exists := g.adjacencies[vertex]; !exists {
		g.adjacencies[vertex] = []Edge{}
	}
}

// AdicionarAresta cria uma conexão entre dois vértices.
// Como nosso grafo é não-direcionado, a conexão é mútua.
func (g *Graph) AddEdge(origin, destiny, weight int) {
	// Garante que ambos os vértices existam no mapa antes de conectar.
	g.AddVertex(origin)
	g.AddVertex(destiny)

	// Adiciona a conexão nos dois sentidos.
	g.adjacencies[origin] = append(g.adjacencies[origin], Edge{Destination: destiny, Weight: weight})
	g.adjacencies[destiny] = append(g.adjacencies[destiny], Edge{Destination: destiny, Weight: weight})
	fmt.Printf("Addicionando aresta: %d <--> %d\n", origin, destiny)
}

// Imprimir mostra a estrutura do grafo (a lista de adjacência).
func (g *Graph) PrintGraph() {
	fmt.Println("\nEstrutura do Grafo Ponderado:")
	for vertex, edges := range g.adjacencies {
		fmt.Printf("Vértice %d -> ", vertex)
		for _, aresta := range edges {
			fmt.Printf("(Dest: %d, Peso: %d) ", aresta.Destination, aresta.Weight)
		}
		fmt.Println()
	}
}

// BFS realiza uma busca em largura a partir de um vértice inicial.
// Ele imprime a ordem em que os vértices são visitados.
func (g *Graph) BFS(begin int) {

	// A fila de vértices a serem visitados. Começa com o nó inicial.
	// Usaremos um slice como uma fila simples para este exemplo.
	queue := []int{begin}

	// O mapa de visitados para evitar ciclos e trabalho repetido.
	// A chave é o vértice (int) e o valor é um booleano.
	visited := make(map[int]bool)
	visited[begin] = true

	fmt.Printf("\nIniciando BFS a partir do vertice %d...\nOrdem de visita: ", begin)

	for len(queue) > 0 {
		// 1. Tira o primeiro elemento da fila (Dequeue)
		currentVertex := queue[0]
		queue = queue[1:] // Encurta a fila

		fmt.Printf("%d ", currentVertex)

		// 2. Para cada vizinho do vértice atual...
		for _, edge := range g.adjacencies[currentVertex] {
			// 3. ...se o vizinho ainda NÃO foi visitado...
			neighbor := edge.Destination
			if !visited[neighbor] {
				// 4. ...marca como visitado e adiciona ao fim da fila.
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}

		}
	}
	fmt.Println("\nBFS concluida.")
}

// DFS é o método público que inicia a busca em profundidade.
func (g *Graph) DFS(begin int) {
	// O mapa de visitados é compartilhado por todas as chamadas recursivas.
	visited := make(map[int]bool)

	fmt.Printf("\nIniciando DFS a partir de %d...\nOrdem de visita: ", begin)

	// Inicia a recursão
	g.dfsRecusive(begin, visited)

	fmt.Println("\nDFS concluido")
}

// dfsRecursivo é a função auxiliar que faz o trabalho pesado.
func (g *Graph) dfsRecusive(currentVertex int, visited map[int]bool) {
	// 1. Marca o nó atual como visitado e o processa.
	visited[currentVertex] = true
	fmt.Printf("%d ", currentVertex)
	// 2. Para cada vizinho do nó atual...
	for _, edge := range g.adjacencies[currentVertex] {
		neighbor := edge.Destination
		// 3. ...se o vizinho ainda NÃO foi visitado...
		if !visited[neighbor] {
			// 4. ...mergulha mais fundo na recursão a partir dele.
			g.dfsRecusive(neighbor, visited)
		}
	}
	// Quando o 'for' termina, significa que batemos num "beco sem saída".
	// A função retorna, e a execução volta para a chamada anterior (backtracking).
}

func (g *Graph) FindVertex(begin, target int) bool {

	queue := []int{begin}

	visited := make(map[int]bool)
	visited[begin] = true

	fmt.Printf("\nIniciando BFS a partir do vertice %d...\nOrdem de visita: ", begin)

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", currentVertex)

		if currentVertex == target {
			// Achou o vertice
			fmt.Printf("\nBFS concluida. Vertice %d encontrado", target)
			return true
		}

		// 2. Para cada vizinho do vértice atual...
		for _, edge := range g.adjacencies[currentVertex] {
			neighbor := edge.Destination
			// 3. ...se o vizinho ainda NÃO foi visitado...
			if !visited[neighbor] {
				// 4. ...marca como visitado e adiciona ao fim da fila.

				visited[neighbor] = true
				queue = append(queue, neighbor)
			}

		}
	}
	fmt.Printf("\nBFS concluida. Vertice %d não encontrado", target)
	// Não achou o vertice
	return false
}

func (g *Graph) Dijkstra(start int) (map[int]int, map[int]int) {
	// distancias: armazena o custo do caminho mais curto da partida até cada vértice
	distancies := make(map[int]int)
	// anteriores: ajuda a reconstruir o caminho mais curto
	previous := make(map[int]int)
	// Fila de prioridade para os vértices a serem visitados
	pq := make(PriorityQueue, 0)

	// Inicialização
	for v := range g.adjacencies {
		distancies[v] = math.MaxInt32
		previous[v] = -1 // -1 indica nenhum anterior
	}
	distancies[start] = 0

	heap.Init(&pq)
	heap.Push(&pq, &Item{Vertex: start, Weight: 0})

	for pq.Len() > 0 {
		// Pega o vertice com a menor distancia da partida
		item := heap.Pop(&pq).(*Item)
		u := item.Vertex

		// Otimização: se já encontramos um caminho mais curto para 'u', pulamos
		if item.Weight > distancies[u] {
			continue
		}

		// Para cada vizinho 'v' de 'u'
		for _, edge := range g.adjacencies[u] {
			v := edge.Destination
			weightUV := edge.Weight

			// Se encontramos um caminho mais curto para 'v' passando por 'u'
			if distancies[u]+weightUV < distancies[v] {
				// Atualizamos a distância e o caminho
				distancies[v] = distancies[u] + weightUV
				previous[v] = u
				// Adicionamos 'v' à fila com sua nova (melhor) distância
				heap.Push(&pq, &Item{Vertex: v, Weight: distancies[v]})
			}
		}
	}
	return distancies, previous
}

// Função auxiliar para imprimir o caminho
func (g *Graph) PrintPath(start, destiny int, previous map[int]int)  {
	path := []int{}
	current := destiny

	for current != -1 {
		path = append([]int{current}, path...)
		current = previous[current]
	}
	if len(path) > 0 && path[0] == start{
		fmt.Printf("Caminho: %v\n", path)

	} else {
		fmt.Printf("Nenhum caminho encontrado de %d para %d\n", start, destiny)

	}
}
// Algoritmo de Prim
func (g *Graph) PrimMST(start int) {
	fmt.Println("\n--- Iniciando Algoritmo de Prim a partir do vértice", start, "---")

	// Mapa para guardar os pesos mínimos para alcançar cada vértice (inicialmente infinito)
	weights := make(map[int]int)
	// Mapa para rastrear qual aresta nos levou a cada vértice na MST
	edgesFromMST := make(map[int]Edge)
	// Mapa para saber quais vértices já estão na MST
	onMST := make(map[int]bool)

	for v := range g.adjacencies {
		weights[v] = math.MaxInt32
	}

	weights[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Vertex: start, Weight: 0})

	totalCost := 0

	for pq.Len() > 0 {
		// Pega o vértice com a aresta de menor peso que conecta à MST
		item := heap.Pop(&pq).(*Item)
		u := item.Vertex

		// Se já incluímos este vértice, pulamos para evitar ciclos
		if onMST[u] {
			continue
		}
		// Adiciona o vértice e a aresta à MST
		onMST[u] = true
		totalCost += weights[u]
		if u != start {
			edge := edgesFromMST[u]
			fmt.Printf("Adicionando aresta: %d --(%d)--> %d\n", edge.Destination, edge.Weight, u)

		}

		// Relaxa as arestas dos vizinhos de u
		for _, neighborEdge := range g.adjacencies[u] {
			v := neighborEdge.Destination
			weightUV := neighborEdge.Weight

			// Se 'v' não está na MST e encontramos um caminho mais barato para ele
			if !onMST[v] && weightUV < weights[v] {
				weights[v] = weightUV
				edgesFromMST[v] = Edge{Destination: u, Weight: weightUV}
				heap.Push(&pq, &Item{Vertex: v, Weight: weightUV})
			}
		}
	}
	fmt.Printf("\nCusto Total da Árvore Geradora Mínima: %d\n", totalCost)
}

// --- Função Principal para Criar e Visualizar um Grafo ---
func main() {
	meuGrafo := NewGraph()

	// Adicionando vértices e arestas para criar um grafo.
	// Vamos imaginar uma pequena rede social.
	meuGrafo.AddEdge(1, 2, 7) // 1 e 2 são amigos
	meuGrafo.AddEdge(1, 3, 9) // 1 e 3 são amigos
	meuGrafo.AddEdge(2, 4, 10)
	meuGrafo.AddEdge(3, 5, 15)
	meuGrafo.AddEdge(4, 5, 11)
	meuGrafo.AddEdge(5, 6, 5)

	// Adicionando um vértice isolado, sem conexões.
	meuGrafo.AddVertex(6)

	// Visualizando o resultado.
	meuGrafo.PrintGraph()
	meuGrafo.BFS(1)
	meuGrafo.FindVertex(1, 9)

	meuGrafo.DFS(1)

	meuGrafo.PrimMST(1)


	partida := 1
	distancias, anteriores := meuGrafo.Dijkstra(partida)

	fmt.Println("--- Algoritmo de Dijkstra a partir do vértice", partida, "---")
	for vertice, dist := range distancias {
		fmt.Printf("Distância mais curta para %d: %d\n", vertice, dist)
	}
	
	fmt.Println("\n--- Reconstruindo o caminho mais curto ---")
	chegada := 4
	fmt.Printf("Caminho mais curto de %d para %d:\n", partida, chegada)
	meuGrafo.PrintPath(partida, chegada, anteriores)
}
