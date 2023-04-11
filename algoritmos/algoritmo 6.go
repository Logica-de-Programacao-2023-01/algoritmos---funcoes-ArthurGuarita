//Escreva uma função que receba um slice de 'strings' como parâmetro e retorne uma 'string' com todas
//as 'strings' concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func conc(s []string) string {
	if len(s) == 0 {
		return string(-1)
	}
	str := strings.Join(s, ", ")
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
	fmt.Println("Strings digitadas:", conc(s))
}
