package main

import "fmt"

func main() {
	var s []int
	var n, aux int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	for i := 0; i <= len(s); i++ {
		fmt.Println("Digite um número para o slice (digite 0 para sair): ")
		fmt.Scan(&aux)
		if aux == 0 {
			break
		}
		s = append(s, aux)
	}
	result, err := present(n, s)
	if err != nil {
		fmt.Println("erro, slice está vazio")
	} else {
		fmt.Println(result)
	}
}
func present(n int, s []int) (bool, error) {
	if len(s) == 0 {
		return false, fmt.Errorf("ocorreu um erro")
	}
	for _, num := range s {
		if num == n {
			return true, nil
		}
	}
	return false, nil
}
