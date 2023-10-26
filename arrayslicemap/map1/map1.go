package main

import "fmt"

func main() {
    // var aprovados map[int]string
    // Mapas devem ser inicializados
    aprovados := make(map[int]string)

    aprovados[12345678978] = "Maria"
    aprovados[98765432123] = "Pedro"
    aprovados[95135745692] = "Ana"

    fmt.Println(aprovados)


    for cpf, nome := range aprovados {
        fmt.Printf("%s (CPF: %d)\n", nome, cpf)
    }

    fmt.Println(aprovados[12345678978])
    delete(aprovados, 12345678978)
    fmt.Println(aprovados[12345678978])
}
