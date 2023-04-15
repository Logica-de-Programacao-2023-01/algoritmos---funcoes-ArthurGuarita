package main

import (
	"errors"
	"fmt"
)

// Função para filtrar strings com mais de 5 caracteres
func filtrarStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	var resultado []string
	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}
	return resultado, nil
}

func main() {
	// Exemplo de uso da função filtrarStrings
	strings := []string{"abc", "defgh", "ijklmno", "pqr", "stuv", "wxyz"}
	resultado, err := filtrarStrings(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings filtradas:", resultado)
	}
}
