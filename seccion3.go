package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Creación de Estructuras (colección de campos)

type Clientes struct { //se define el cliente y el atributo
	Clientes chan int
}

type Cajeros struct { //se definen los cajeros y el atributo
	Cajeros chan int
}

//Creación de Funciones y metodos

//Función de Clientes

func (cl *Clientes) rutina(tiempo int) { //Metodo rutina
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

//Función de Cajeros (Generadora)

func (c *Cajeros) rutina(cl *Clientes, tiempo int) { //Metodo rutina
	i := 0

	for {
		if cap(c.Cajeros) == len(c.Cajeros) { //si la capacidad es igual al tamaño ingresado o llenado
			//fmt.Println("No hay caja")
			//time.Sleep(1 * time.Second)
			continue
		} else { //si no está llena, ingresa un cliente a la fila con el tiempo asignado que es variable
			rand.Seed(time.Now().UnixNano())
			delta := rand.Intn(2)
			time.Sleep(time.Duration(delta) * time.Second)
			i = <-cl.Clientes
			c.Cajeros <- i
			go tiempoCaja(c.Cajeros, i, tiempo)
		}
	}
}

//Función del Tiempo en Caja

func tiempoCaja(Cajeros chan int, n int, tiempo int) { //Tiempo de un cliente en caja
	rand.Seed(time.Now().UnixNano())
	delta := rand.Intn(tiempo)
	fmt.Printf("Cliente %d en caja\n", n)
	time.Sleep(time.Duration(delta) * time.Second)
	fmt.Printf("Cliente %d salio de caja.\n", <-Cajeros)
	time.Sleep(1 * time.Second)
}

func main() {

	//Creación
	cajeros := new(Cajeros)
	clientes := new(Clientes)

	//Asignación del Cajero
	cantidadCajero := 0
	fmt.Printf("Ingrese cantidad de Cajeros: ")
	fmt.Scanf("%d\n", &cantidadCajero)

	//Asignación del Cajero
	cantidadFila := 0
	fmt.Printf("Ingrese cantidad de filas: ")
	fmt.Scanf("%d\n", &cantidadFila)

	//Asignación del campo de la estructura
	cajeros.Cajeros = make(chan int, cantidadCajero)
	clientes.Clientes = make(chan int, cantidadFila)

	//Asignación del tiempo (utilizado para entrada y salida de clientes de forma variable)
	tiempo := 0
	fmt.Printf("Ingrese probabilidad de tiempo: ")
	fmt.Scanf("%d\n", &tiempo)

	//Llamado mediante Gorrutina
	//Acceder a los valores a través de la notación de puntos y parametros
	go clientes.rutina(tiempo)
	go cajeros.rutina(clientes, tiempo)

	//aqui se sale al recibir algo por consola (fin del programa) -> puede ser tecla Enter
	waiter := ""
	fmt.Scanf("%s\n", &waiter)

}
