package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	result, err := vogais(str)
	if err != nil {
		fmt.Print("Ocorreu um erro")
	} else {
		fmt.Println(result)
	}
}
func vogais(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("ocorreu um erro")
	}
	var vogal = "AEIOUaeiou"
	for _, c := range vogal {
		str = strings.ReplaceAll(str, string(c), "*")
	}
	return str, nil
}
