package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- "Capuccino"
		c <- "Latte"
		c <- "No"
		fmt.Println("Cliente: Tenga un buen día")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("\nBienvenido a Code Cofee :)\n\nVendedor: Que café desea pedir?")
		time.Sleep(time.Second * 2)
		pedido := <-c //bloqueante
		fmt.Println("Cliente: " + pedido)
		time.Sleep(time.Second * 2)
		pedido2 := <-c //bloqueante
		fmt.Println("Cliente: Mejor un " + pedido2 + "!")
		time.Sleep(time.Second * 2)
		fmt.Println("Vendedor: OK, se cambia entonces será un", pedido2)
		time.Sleep(time.Second * 1)
		fmt.Println("Vendedor: Desea algo más?")
		time.Sleep(time.Second * 2)
		respuesta := <-c //bloqueante
		fmt.Println("Cliente: ", respuesta)
	}()

	wg.Wait()
	close(c)
}
