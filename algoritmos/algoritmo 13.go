package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	result, err := soma(n)
	if err != nil {
		fmt.Println("erro")
	} else {
		fmt.Println(result)
	}
}

func soma(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("o número é negativo")
	}
	sum := 0
	v := strconv.Itoa(n)
	for _, c := range v {
		alg, _ := strconv.Atoi(string(c))
		sum += alg
	}
	return sum, nil
}
