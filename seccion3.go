package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("---> Ingresa Cliente: ", i)
			fmt.Println("Cantidad:", cap(ch), " utilizado:", len(ch))
		}
		fmt.Println("xxxxxxx Cajas Saturadas xxxxxxx")
		close(ch)
	}()
	for valor := range ch {
		fmt.Println("---> Termina Cliente: ", valor)
		fmt.Println("Cantidad:", cap(ch), " utilizado:", len(ch))
	}
	fmt.Println("Termina")
}
