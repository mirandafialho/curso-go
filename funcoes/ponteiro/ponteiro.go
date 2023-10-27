package main

import "fmt"

func inc1(n int) {
	n++
}

// Revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// Revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor (cópia)
	fmt.Println(n)

	// Revisão: & é usado para obter o endereço da variável
	inc2(&n)
	fmt.Println(n)
}
