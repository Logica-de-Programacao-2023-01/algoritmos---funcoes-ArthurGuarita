package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)
	//
	fmt.Print(vogais(str))
}
func vogais(str string) int {
	vogal := "AEIOUaeiou"
	count := 0
	for _, v := range str {
		if strings.ContainsRune(vogal, v) {
			count++
		}
	}
	return count
}
