package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 20))
}

func Soma(a int, b int) int {
	return a + b
}

func Divisao(a int, b int) int {

	if b == 0 {
	 return 0
	}
	
	return a / b
}


