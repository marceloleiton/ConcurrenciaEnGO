package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clientes struct {
	Clientes chan int
}

func (cl *Clientes) rutina(tiempo int) {
	i := 1
	for {
		if cap(cl.Clientes) == len(cl.Clientes) {
			//fmt.Println("No hay cupos Clientes")
			//time.Sleep(1 * time.Second)
			continue
		} else {
			rand.Seed(time.Now().UnixNano())
			delta := rand.Intn(tiempo)
			cl.Clientes <- i
			i++
			time.Sleep(time.Duration(delta) * time.Second)
		}
	}
}

type Cajeros struct {
	Cajeros chan int
}

func (c *Cajeros) rutina(cl *Clientes, tiempo int) {
	i := 0

	for {
		if cap(c.Cajeros) == len(c.Cajeros) {
			//fmt.Println("No hay caja")
			//time.Sleep(1 * time.Second)
			continue
		} else {
			rand.Seed(time.Now().UnixNano())
			delta := rand.Intn(2)
			time.Sleep(time.Duration(delta) * time.Second)
			i = <-cl.Clientes
			c.Cajeros <- i

			go tiempoCaja(c.Cajeros, i, tiempo)
		}
	}
}

func tiempoCaja(Cajeros chan int, n int, tiempo int) {
	rand.Seed(time.Now().UnixNano())
	delta := rand.Intn(tiempo)
	fmt.Printf("Cliente %d en caja\n", n)
	time.Sleep(time.Duration(delta) * time.Second)
	fmt.Printf("Cliente %d salio de caja.\n", <-Cajeros)
	time.Sleep(1 * time.Second)

}

func main() {

	cajeros := new(Cajeros)
	clientes := new(Clientes)
	cajeros.Cajeros = make(chan int, 3)
	clientes.Clientes = make(chan int, 5)
	tiempo := 0
	fmt.Printf("Ingrese tiempo: ")
	fmt.Scanf("%d\n", &tiempo)

	go clientes.rutina(tiempo)
	go cajeros.rutina(clientes, tiempo)

	//aqui se sale al recibir algo por consola (fin del programa)
	waiter := ""
	fmt.Scanf("%s\n", &waiter)

}
