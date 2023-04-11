package main

import "fmt"

func menorValue(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("A lista de números está vazia. ")
	}
	min := list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min, nil
}
func main() {
	var list []int
	var num int
	for i := 0; i <= len(list); i++ {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		list = append(list, num)
	}
	result, err := menorValue(list)
	if err != nil {
		fmt.Print("Ocorreu um erro")
	}
	fmt.Print(result)
}
