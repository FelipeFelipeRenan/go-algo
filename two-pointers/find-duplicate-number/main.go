package main

import "fmt"

func main() {
	fmt.Println("--- Find the Duplicate Number ---")

	exemplo := []int{1, 3, 4, 2, 2}

	fmt.Printf("Array de entrada: %v\n", exemplo)
	fmt.Printf("Número duplicado encontrado: %d\n", FindDuplicate(exemplo))
}
