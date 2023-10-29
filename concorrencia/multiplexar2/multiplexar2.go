package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	canal := make(chan string)
	go func() {
		for index := 0; index < 3; index++ {
			time.Sleep(time.Second)
			canal <- fmt.Sprintf("%s falando: %d", pessoa, index)
		}
	}()

	return canal
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			select {
			case saida := <-entrada1:
				canal <- saida
			case saida := <-entrada2:
				canal <- saida
			}
		}
	}()

	return canal
}

func main() {
	canal := juntar(falar("Joao"), falar("Maria"))

	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
	fmt.Println(<-canal, <-canal)
}
