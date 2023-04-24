package main

import (
	"errors"
	"fmt"
)

// Função para verificar se um número é primo
func isPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

// Função para obter os números primos menores ou iguais a um número
func numerosPrimosAteN(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("o número deve ser maior ou igual a zero")
	}

	var primos []int
	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos, nil
}

func main() {
	// Exemplo de uso da função numerosPrimosAteN
	numero := 20
	primos, err := numerosPrimosAteN(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos até", numero, ":", primos)
	}
}
