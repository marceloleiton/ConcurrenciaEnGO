package main

import (
	"fmt"
	"time"
)

//Creación de Estructuras (colección de campos)

type Producto struct {
	chA chan string
}

type Consumidor struct {
	chB chan string
}

//Creación de Funciones y metodos

//Función del Productor
func (p *Producto) rutina(c *Consumidor) { //Metodo rutina
	fmt.Println("Begin A ↓")
	p.chA <- "Call B (corroutine) ↗" //se envia a corrutina B
	fmt.Println(<-c.chB)             //Resume A
	fmt.Println("continue in A")
	p.chA <- "resume B"
	fmt.Println(<-c.chB) //end B
	fmt.Println("End A")
}

//Función del Consumidor
func (c *Consumidor) rutina(p *Producto) { //Metodo rutina
	fmt.Println(<-p.chA) //Call B (corroutine) ↗
	fmt.Println("Begin B ↓")
	fmt.Println("continue in B")
	c.chB <- "Resume A ←"
	fmt.Println(<-p.chA) // Resume B
	fmt.Println("continue in B")
	c.chB <- "return end B"
}

func main() {

	//Creación
	consumidor := new(Consumidor)
	productor := new(Producto)
	//Asignación del campo de la estructura
	productor.chA = make(chan string)
	consumidor.chB = make(chan string)

	//Llamado mediante Gorrutina
	//Acceder a los valores a través de la notación de puntos
	go productor.rutina(consumidor)
	go consumidor.rutina(productor)

	time.Sleep(5 * time.Second)
}
