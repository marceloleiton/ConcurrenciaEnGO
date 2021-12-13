package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("Cajeros:", cap(ch), " utilizado:", len(ch))
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("Cliente: ", i, "Ingresa")
			fmt.Println("Cajeros:", cap(ch), " utilizado:", len(ch))
		}
		fmt.Println("No hay cupos")
	}()

	time.Sleep(1 * time.Second)

	for valor := range ch {
		fmt.Println("Termina su turno el usuario: ", valor)
		fmt.Println("Cajeros:", cap(ch), " utilizado:", len(ch))
	}
	fmt.Println("saliendo")
	fmt.Println("Cajeros:", cap(ch), " utilizado:", len(ch))
}
