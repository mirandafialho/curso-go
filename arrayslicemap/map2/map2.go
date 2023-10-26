package main

import "fmt"

func main() {
    funcionarioSalario := map[string]float64 {
        "José João": 11325.45,
        "Gabriela Silva": 15456.78,
        "Pedro Junior": 1200.00,
    }

    funcionarioSalario["Rafael Filho"] = 1350.00
    delete(funcionarioSalario, "Inexistente")

    for nome, salario := range funcionarioSalario {
        fmt.Println(nome, salario)
    }
}
