package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()
	result, err := words(str)
	if err != nil {
		fmt.Errorf("Ocorreu um erro.")
	}
	fmt.Println(result)

}
func words(str string) ([]string, error) {
	var word []string
	if str == "" {
		return nil, errors.New("não pode ser vazia")
	}
	word = strings.Fields(str)
	return word, nil
}
