package main

import (
	"fmt"
	"strings"
)

func concatenar(s []string) string {
	str := strings.Join(s, " ")
	return str
}

func main() {
	var s []string
	var v string
	for {
		fmt.Print("Digite uma string (ou 'q' para sair): ")
		fmt.Scan(&v)
		if v == "q" {
			break
		}
		s = append(s, v)
	}
	fmt.Println("Strings digitadas:", concatenar(s))
}
