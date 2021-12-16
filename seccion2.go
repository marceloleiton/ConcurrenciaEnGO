package main

import (
	"fmt"
	"time"
)

type Producto struct { //se define el producto y los atributos que va a tener
	chA chan string
}

func (p *Producto) rutina(c *Consumidor) { //corrutina A
	fmt.Println("Begin A ↓")         //primer print
	p.chA <- "Call B (corroutine) ↗" //se envia a corrutina B

	fmt.Println(<-c.chB) //Resume A
	fmt.Println("continue in A")
	p.chA <- "resume B"
	fmt.Println(<-c.chB) //end B
	fmt.Println("End A")
}

type Consumidor struct { //se define el consumidor y el atributo
	chB chan string
}

func (c *Consumidor) rutina(p *Producto) { //corrutina B
	fmt.Println(<-p.chA) //Call B (corroutine) ↗
	fmt.Println("Begin B ↓")
	fmt.Println("continue in B")
	c.chB <- "Resume A ←"
	fmt.Println(<-p.chA) // Resume B
	fmt.Println("continue in B")
	c.chB <- "return end B"
}

func main() {

	consumidor := new(Consumidor)      //seteamos un dato de la estructura
	productor := new(Producto)         //seteamos un dato de la estructura
	productor.chA = make(chan string)  //asignamos el dato
	consumidor.chB = make(chan string) //asignamos el dato

	go productor.rutina(consumidor)
	go consumidor.rutina(productor)

	time.Sleep(5 * time.Second)
}
