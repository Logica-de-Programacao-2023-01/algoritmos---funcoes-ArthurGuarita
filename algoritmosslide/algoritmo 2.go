package main

import "fmt"

func main() {
	var list []int
	var num int

	for i := 0; i <= len(list); i++ {
		fmt.Print("Digite os valores da lista (termine com 0): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		list = append(list, num)
	}
	result, err := media(list)
	if err != nil {
		fmt.Print("OCorreu um erro")
	}
	fmt.Print(result)

}
func media(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("A lista está vazia. ")
	}
	//
	var sum int
	for _, soma := range list {
		sum += soma
	}
	var media int
	media = sum / len(list)
	return media, nil
}
