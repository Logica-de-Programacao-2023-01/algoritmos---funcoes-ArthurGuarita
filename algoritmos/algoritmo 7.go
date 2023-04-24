package main

import (
	"errors"
	"fmt"
)

// Função genérica que recebe um slice de inteiros e uma função a ser aplicada
// a cada elemento do slice
func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// Definindo uma função de exemplo que multiplica cada elemento por 2
	funcaoExemplo := func(n int) int {
		return n * 2
	}

	// Chamando a função aplicarFuncao com o slice e a função de exemplo
	resultado, err := aplicarFuncao(slice, funcaoExemplo)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado)
}
