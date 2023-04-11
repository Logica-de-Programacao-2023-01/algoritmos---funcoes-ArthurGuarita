package main

import "fmt"

func main() {
	var f []int
	for i := 0; i <= 10; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	fmt.Print(f)
}
