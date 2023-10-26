package main

import "fmt"

func main() {
    funcionariosPorLetra := map[string]map[string]float64 {
        "G": {
            "Gabriela Silva": 15456.78,
            "Guga Pereira": 8456.78,
        },
        "J": {
            "Jose Joao": 11324.45,
        },
        "P": {
            "Pedro Junior": 1200.00,
        },
    }


    delete(funcionariosPorLetra, "P")
    
    for letra, funcionarios := range funcionariosPorLetra {
        fmt.Println(letra, funcionarios)
    }
}
