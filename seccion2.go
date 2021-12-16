package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		fmt.Println(<-ch2)
		time.Sleep(5 * time.Second)
		ch <- "Capuccino :)"
		fmt.Println(<-ch2)
		ch <- "Terminado ."
	}()

	fmt.Println("Vendedor: Que café desea pedir?")
	ch2 <- "Toma un tiempo para pensar"
	valor := <-ch //Bloqueante

	fmt.Println(valor)
	time.Sleep(5 * time.Second)
	ch2 <- "Mejor un Latte !"
	fmt.Println("Vendedor: OK, se cambia entonces :)", <-ch)
}
