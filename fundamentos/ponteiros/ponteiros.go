package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // Pegar endereço
	*p++   // Desreferenciando (pegando o valor)

	// Go não tem aritmética de ponteiros.
	// p++

	fmt.Println(p, *p, i, &i)
}
