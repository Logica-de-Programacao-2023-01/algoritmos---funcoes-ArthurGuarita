package main

import "fmt"

// Crie uma função que receba um número inteiro como parâmetro e retorne a sequência
// de Fibonacci correspondente a esse número. Caso o número seja negativo, retorne um
// erro.
func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("O valor digitado é negativo")
	}
	//sequencia
	var f []int
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	last := f[n-1]
	return last, nil

}

func main() {
	var n int
	fmt.Print("Digite um valor inteiro: ")
	fmt.Scan(&n)

}
