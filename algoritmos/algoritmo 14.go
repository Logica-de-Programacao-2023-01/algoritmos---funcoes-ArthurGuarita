package main

import "fmt"

func main() {
	var a, b []int
	var a1, b1 int
	for {
		fmt.Println("Digite um número para o slice 'a' (0 para sair): ")
		fmt.Scan(&a1)
		if a1 == 0 {
			break
		}
		a = append(a, a1)
	}
	for {
		fmt.Println("Digite um valor para o slice 'b' (0 para sair): ")
		fmt.Scan(&b1)
		if b1 == 0 {
			break
		}
		b = append(b, b1)
	}
	result, err := inter(a, b)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		fmt.Println(result)
	}
}

func inter(a []int, b []int) ([]int, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("ocorreu um erro no slice a")
	} else if len(b) == 0 {
		return nil, fmt.Errorf("ocorreu um erro no slice b")
	}

	var equal []int
	if len(a) > len(b) {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					equal = append(equal, a[i])
				}
			}
		}
	} else {
		for i := 0; i < len(b); i++ {
			for j := 0; j < len(a); j++ {
				if b[i] == a[j] {
					equal = append(equal, b[i])
				}
			}
		}
	}

	return equal, nil
}
