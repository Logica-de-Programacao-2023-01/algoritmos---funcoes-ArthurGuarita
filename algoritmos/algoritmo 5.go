package main

//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro
//elemento igual ao valor no slice. Caso não encontre, retorne -1.

func encontrarIndex(s []int, n int) int {
	for i, valor := range s {
		if s[i] == n {
			return i
		}
	}
	return -1
}
