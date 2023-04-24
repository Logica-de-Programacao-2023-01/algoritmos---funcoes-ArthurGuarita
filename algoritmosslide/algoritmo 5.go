package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o valor de 'n' para a sequência de Fibonacci: ")
	fmt.Scan(&n)

	result, err := f(n)
	if err != nil {
		fmt.Printf("Ocorreu um erro: %v\n", err)
		return
	}
	fmt.Println(result)
}

func f(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("O valor de 'n' deve ser maior ou igual a 0")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}

	s := make([]int, n+1)
	s[0], s[1] = 0, 1

	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	last := s[n]
	return last, nil
}
