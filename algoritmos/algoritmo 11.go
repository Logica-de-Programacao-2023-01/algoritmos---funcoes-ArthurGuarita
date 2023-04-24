package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	result, err := EPrimo(n)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		fmt.Println(result)
	}
}

func EPrimo(n int) (bool, error) {
	if n <= 1 {
		return false, fmt.Errorf("o número deve ser maior que 1")
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
