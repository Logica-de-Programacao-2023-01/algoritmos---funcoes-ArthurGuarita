package main

import (
	"fmt"
	"sort"
)

func crescente(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, fmt.Errorf("a lista está vazia")
	}

	// Utiliza a função sort.Ints para ordenar os números em ordem crescente
	sort.Ints(nums)
	return nums, nil
}

func main() {
	var nums []int
	for {
		var n int
		fmt.Println("Digite um número (digite 0 para sair): ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		nums = append(nums, n)
	}
	ordem, err := crescente(nums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Números em ordem crescente:", ordem)
}
