package main

import (
	"fmt"
)

func numCarac(str string) (int, error) {
	if len(str) == 0 {
		return 0, fmt.Errorf("Não foi digitado nenhuma string.")
	}
	count := 0
	for i := 0; i <= len(str); i++ {
		count = i
	}
	return count, nil
}
func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	result, err := numCarac(str)
	if err != nil {
		fmt.Errorf("Não possui nenhuma caractere. ")
	} else {
		fmt.Print(result)
	}
}
