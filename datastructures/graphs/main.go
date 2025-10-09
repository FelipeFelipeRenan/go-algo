package main

import "fmt"

// Grafo representa a estrutura de um grafo usando uma lista de adjacência.
type Graph struct {
	// O mapa onde a chave é um vértice (int) e o valor é um slice
	// de todos os vértices vizinhos a ele.
	adjacencies map[int][]int
}

// NovoGrafo cria e inicializa um grafo vazio.
func NewGraph() *Graph {
	return &Graph{
		adjacencies: make(map[int][]int),
	}
}

// AdicionarVertice garante que um vértice exista no mapa, mesmo que sem arestas.
func (g *Graph) AddVertex(vertex int) {
	// Se o vértice ainda não estiver no mapa, adiciona com uma lista vazia.
	if _, exists := g.adjacencies[vertex]; !exists {
		g.adjacencies[vertex] = []int{}
	}
}

// AdicionarAresta cria uma conexão entre dois vértices.
// Como nosso grafo é não-direcionado, a conexão é mútua.
func (g *Graph) AddEdge(vertex1, vertex2 int) {
	// Garante que ambos os vértices existam no mapa antes de conectar.
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)

	// Adiciona a conexão nos dois sentidos.
	g.adjacencies[vertex1] = append(g.adjacencies[vertex1], vertex2)
	g.adjacencies[vertex2] = append(g.adjacencies[vertex2], vertex1)
	fmt.Printf("Addicionando aresta: %d <--> %d\n", vertex1, vertex2)
}

// Imprimir mostra a estrutura do grafo (a lista de adjacência).
func (g *Graph) PrintGraph() {
	fmt.Println("\nEstrutura do Grafo (Lista de adjacencia):")
	for vertex, adjacents := range g.adjacencies {
		fmt.Printf("Vertice %d -> Vizinhos: %v\n", vertex, adjacents)
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
		for _, neighbor := range g.adjacencies[currentVertex] {
			// 3. ...se o vizinho ainda NÃO foi visitado...
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
	for _, neighbor := range g.adjacencies[currentVertex] {
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
		for _, neighbor := range g.adjacencies[currentVertex] {
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

// --- Função Principal para Criar e Visualizar um Grafo ---
func main() {
	meuGrafo := NewGraph()

	// Adicionando vértices e arestas para criar um grafo.
	// Vamos imaginar uma pequena rede social.
	meuGrafo.AddEdge(1, 2) // 1 e 2 são amigos
	meuGrafo.AddEdge(1, 3) // 1 e 3 são amigos
	meuGrafo.AddEdge(2, 4)
	meuGrafo.AddEdge(3, 5)
	meuGrafo.AddEdge(4, 5)
	meuGrafo.AddEdge(5, 6)

	// Adicionando um vértice isolado, sem conexões.
	meuGrafo.AddVertex(6)

	// Visualizando o resultado.
	meuGrafo.PrintGraph()
	meuGrafo.BFS(1)
	meuGrafo.FindVertex(1, 9)

	meuGrafo.DFS(1)

}
