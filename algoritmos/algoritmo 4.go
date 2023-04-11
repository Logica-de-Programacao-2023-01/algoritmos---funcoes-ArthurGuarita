package main

import "fmt"

// Função para encontrar o segundo maior valor em um slice de inteiros
func secondValue(s []int) int {
	if len(s) < 2 {
		// Verifica se o slice tem pelo menos 2 elementos
		// Se não tiver, retorna um valor inválido
		return -1
	}

	max := s[0]
	secondMax := s[1]

	// Encontra o maior valor
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	// Encontra o segundo maior valor
	for _, v := range s {
		if v > secondMax && v < max {
			secondMax = v
		}
	}

	return secondMax
}

func main() {
	var s []int
	var num int
	for {
		fmt.Print("Digite um número (ou 0 para sair): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		s = append(s, num)
	}
	fmt.Println("O segundo maior valor é:", secondValue(s))
}
