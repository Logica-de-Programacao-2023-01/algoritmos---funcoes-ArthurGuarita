package main

import (
	"errors"
	"fmt"
)

func par(s []int) ([]int, error) {
	if len(s) == 0 {
		return nil, errors.New("o slice está vazio")
	}
	var pares []int
	for _, num := range s {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	return pares, nil
}

func main() {
	var s []int
	var num int
	for {
		fmt.Print("Digite um valor (digite 0 para terminar): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		s = append(s, num)
	}
	pares, err := par(s)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", pares)
	}
}
