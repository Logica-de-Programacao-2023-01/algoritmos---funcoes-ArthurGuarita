package main

import (
	"fmt"
	"strings"
)

func main() {
	var s []string
	var stri string
	for i := 0; i <= len(s); i++ {
		fmt.Print("Digite uma string (exit para sair): ")
		fmt.Scan(&stri)
		if stri == "exit" {
			break
		}
		s = append(s, stri)
	}
	result, err := Upper(s)
	if err != nil {
		fmt.Errorf("ocorreu um erro")
	} else {
		fmt.Println(result)
	}
}
func Upper(s []string) ([]string, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("o slice está vazio")
	}
	var up []string
	for _, v := range s {
		if v == strings.Title(v) {
			up = append(up, v)
		}
	}
	return up, nil

}
