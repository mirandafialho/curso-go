package main

import "fmt"

type imprimivel interface {
    toString() string
}

type pessoa struct {
    nome string
    sobrenome string
}

type produto struct {
    nome string
    preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
    return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
    return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(imp imprimivel) {
    fmt.Println(imp.toString())
}

func main() {
    var coisa imprimivel = pessoa{nome: "Roberto", sobrenome: "Silva"}
    fmt.Println(coisa.toString())
    imprimir(coisa)

    coisa = produto{nome: "Calça Jeans", preco: 79.90}
    fmt.Println(coisa.toString())
    imprimir(coisa)

    p2 := produto{nome: "Calça Jeans", preco: 179.90}
    imprimir(p2)
}
