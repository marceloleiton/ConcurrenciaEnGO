package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clientes struct { //se define el cliente y el atributo
	Clientes chan int
}

func (cl *Clientes) rutina(tiempo int) { //corrutina A
	i := 1
	for {
		if cap(cl.Clientes) == len(cl.Clientes) { //revisa si la fila de clientes esta llena
			//fmt.Println("No hay cupos Clientes")
			//time.Sleep(1 * time.Second)
			continue
		} else { //si no esta llena la fila va llenando de a 1
			rand.Seed(time.Now().UnixNano())
			delta := rand.Intn(tiempo)
			cl.Clientes <- i
			i++
			time.Sleep(time.Duration(delta) * time.Second)
		}
	}
}

type Cajeros struct { //se definen los cajeros y el atributo que va a tener
	Cajeros chan int
}

func (c *Cajeros) rutina(cl *Clientes, tiempo int) { //corrutina B
	i := 0

	for {
		if cap(c.Cajeros) == len(c.Cajeros) { //revisa si los cajeros estan llenos
			//fmt.Println("No hay caja")
			//time.Sleep(1 * time.Second)
			continue
		} else { //si no esta lleno va llenando la fila
			rand.Seed(time.Now().UnixNano())
			delta := rand.Intn(2)
			time.Sleep(time.Duration(delta) * time.Second)
			i = <-cl.Clientes
			c.Cajeros <- i

			go tiempoCaja(c.Cajeros, i, tiempo)
		}
	}
}

func tiempoCaja(Cajeros chan int, n int, tiempo int) { //tiempo randomico que se encuentra dentro de la caja
	rand.Seed(time.Now().UnixNano())
	delta := rand.Intn(tiempo)
	fmt.Printf("Cliente %d en caja\n", n)
	time.Sleep(time.Duration(delta) * time.Second)
	fmt.Printf("Cliente %d salio de caja.\n", <-Cajeros)
	time.Sleep(1 * time.Second)

}

func main() {

	cajeros := new(Cajeros)               //Seteamos un dato de la estrucutura
	clientes := new(Clientes)             //seteamos un dato de la estructura
	cajeros.Cajeros = make(chan int, 3)   //asignamos el dato
	clientes.Clientes = make(chan int, 5) //asignamos el dato
	tiempo := 0
	fmt.Printf("Ingrese tiempo: ")
	fmt.Scanf("%d\n", &tiempo)

	go clientes.rutina(tiempo)
	go cajeros.rutina(clientes, tiempo)

	//aqui se sale al recibir algo por consola (fin del programa)
	waiter := ""
	fmt.Scanf("%s\n", &waiter)

}
