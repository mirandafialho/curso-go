package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for index := 0; index < quantidade; index++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, index + 1)
	}
}

func main() {
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("Joao", "So posso falar depois de você!", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!!", 5)
}
