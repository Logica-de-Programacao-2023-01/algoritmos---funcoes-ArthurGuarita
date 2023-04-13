package main

import (
	"fmt"
)

// Função para encontrar o índice do primeiro elemento igual ao valor no slice
func encontrarIndex(s []int, n int) int {
	for i, v := range s {
		if v == n {
			return i
		}
	}
	return -1
}

func main() {
	var s []int
	var n, num int
	fmt.Println("Digite o valor do índice: ")
	fmt.Scan(&n)
	for {
		fmt.Print("Digite um número (digite 0 para sair): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		s = append(s, num)
	}
	result := encontrarIndex(s, n)
	fmt.Println("Índice do primeiro elemento igual ao valor:", result)
}
