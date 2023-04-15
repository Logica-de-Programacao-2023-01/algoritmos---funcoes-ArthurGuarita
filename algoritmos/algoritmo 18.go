package main

import (
	"errors"
	"fmt"
)

// Função para aplicar uma função em cada elemento do slice e retornar a soma dos resultados
func aplicarFuncaoEmSlice(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("o slice está vazio")
	}

	soma := 0
	for _, valor := range slice {
		resultado := funcao(valor)
		soma += resultado
	}

	return soma, nil
}

func main() {
	// Exemplo de um slice de inteiros
	numeros := []int{1, 2, 3, 4, 5}

	// Função para dobrar um número
	dobrar := func(numero int) int {
		return numero * 2
	}

	// Chamando a função aplicarFuncaoEmSlice com o slice de inteiros e a função de dobrar
	resultado, err := aplicarFuncaoEmSlice(numeros, dobrar)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
