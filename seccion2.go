package main

import (
	"fmt"
	"time"
)

type Producto struct {
	chA chan string
}

func (p *Producto) rutina(c *Consumidor) {
	fmt.Println("Begin A ↓")
	p.chA <- "Call B (corroutine) ↗"

	fmt.Println(<-c.chB) //Resume A
	fmt.Println("continue in A")
	p.chA <- "resume B"
	fmt.Println(<-c.chB) //end B
	fmt.Println("End A")
}

type Consumidor struct {
	chB chan string
}

func (c *Consumidor) rutina(p *Producto) {
	fmt.Println(<-p.chA) //Call B (corroutine) ↗
	fmt.Println("Begin B ↓")
	fmt.Println("continue in B")
	c.chB <- "Resume A ←"
	fmt.Println(<-p.chA) // Resume B
	fmt.Println("continue in B")
	c.chB <- "return end B"
}

func main() {

	consumidor := new(Consumidor)
	productor := new(Producto)
	productor.chA = make(chan string)
	consumidor.chB = make(chan string)

	go productor.rutina(consumidor)
	go consumidor.rutina(productor)

	time.Sleep(5 * time.Second)
}
