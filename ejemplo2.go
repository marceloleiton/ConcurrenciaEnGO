package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tiempoCaja(Cajeros chan int, Cajas chan int, n int, caja int) {
	rand.Seed(time.Now().UnixNano())
	delta := rand.Intn(10)
	fmt.Printf("Cliente %d en caja %d\n", n, caja)
	time.Sleep(time.Duration(delta) * time.Second)
	fmt.Printf("Cliente %d salio de caja %d. Tiempo en caja:  %d\n", <-Cajeros, caja, delta)
	time.Sleep(1 * time.Second)
	Cajas <- caja
}

func main() {
	Cajeros := make(chan int, 3)
	Clientes := make(chan int, 5)
	Cajas := make(chan int, 3)
	Cajas <- 1
	Cajas <- 2
	Cajas <- 3

	go func() {
		i := 1
		for {
			if cap(Clientes) == len(Clientes) {
				//fmt.Println("No hay cupos Clientes")
				//time.Sleep(1 * time.Second)
				continue
			} else {
				rand.Seed(time.Now().UnixNano())
				delta := rand.Intn(10)
				Clientes <- i
				i++
				time.Sleep(time.Duration(delta) * time.Second)
			}
		}

	}()

	go func() {
		i := 0
		caja := 0
		for {
			if cap(Cajeros) == len(Cajeros) {
				//fmt.Println("No hay caja")
				//time.Sleep(1 * time.Second)
				continue
			} else {
				rand.Seed(time.Now().UnixNano())
				delta := rand.Intn(2)
				time.Sleep(time.Duration(delta) * time.Second)
				i = <-Clientes
				Cajeros <- i
				caja = <-Cajas
				go tiempoCaja(Cajeros, Cajas, i, caja)
			}
		}
	}()
	/*for {
		rand.Seed(time.Now().UnixNano())
		delta := rand.Intn(20) + 5
		fmt.Println(delta)
		time.Sleep(1 * time.Second)
	}*/
	waiter := ""
	fmt.Scanf("%s\n", &waiter)

}
