package main

import (
	"fmt"
	"time"
)

func corrutinaB(chA chan string, chB chan string) {
	fmt.Println(<-chA) //Call B (corroutine) ↗
	fmt.Println("Begin B ↓")
	fmt.Println("continue in B")
	chB <- "Resume A ←"
	fmt.Println(<-chA)
	fmt.Println("continue in B")
	chB <- "return end B"
}

func main() {

	chA := make(chan string)
	chB := make(chan string)

	go corrutinaB(chA, chB)

	fmt.Println("Begin A ↓")
	chA <- "Call B (corroutine) ↗"

	fmt.Println(<-chB) //Resume A
	fmt.Println("continue in A")
	chA <- "resume B"
	fmt.Println(<-chB) //end B
	fmt.Println("End A")
	time.Sleep(5 * time.Second)
}
