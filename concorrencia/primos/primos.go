package main

import (
	"fmt"
	"time"
)

func isPrimo(number int) bool {
	for index := 2; index < number; index++ {
		if number % index == 0 {
			return false
		}
	}

	return true
}

func primos(number int, channel chan int) {
	inicio := 2
	for index := 0; index < number; index++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				channel <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(channel)
}

func main () {
	mainChannel := make(chan int, 30)
	go primos(cap(mainChannel), mainChannel)
	
	for primo := range mainChannel {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
