package main

import (
	"fmt"
)

func main() {
	var s []string
	var a string
	for {
		fmt.Print("Digite uma letra (exit para sair): ")
		fmt.Scan(&a)
		if a == "exit" {
			break
		}
		s = append(s, a)
	}
	result, err := ordemAlf(s)
	if err != nil {
		fmt.Print("Erro:", err)
	} else {
		fmt.Println("Ordem alfabética:", result)
	}
}

func ordemAlf(s []string) ([]string, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("o slice está vazio")
	}

	// Copiar o slice original para evitar modificar o original
	ord := make([]string, len(s))
	copy(ord, s)

	// Ordenar o slice em ordem alfabética
	for i := 0; i < len(ord)-1; i++ {
		if ord[i] > ord[i+1] {
			return nil, fmt.Errorf("os elementos não estão em ordem alfabética")
		}
	}

	return ord, nil
}
