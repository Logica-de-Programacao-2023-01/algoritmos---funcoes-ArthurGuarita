package main

import "fmt"

func main() {
	var s []int
	var aux int
	for {
		fmt.Print("Digite um valor (termine a operação com 0): ")
		fmt.Scan(&aux)
		if aux == 0 {
			break
		}
		s = append(s, aux)
	}
	fmt.Printf("Média: %.2f\n", media(s))
}

func media(s []int) float64 {
	soma := 0
	for _, num := range s {
		soma += num
	}
	media := float64(soma) / float64(len(s))
	return media
}
