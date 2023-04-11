package main

import "fmt"

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não pode ser dividido por 0")
	}
	return a / b, nil
}
