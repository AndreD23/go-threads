package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func worker(workerId int, data chan int) {
	// Faz um looping infinito onde ele vai ficar lendo todos os dados que entram neste canal
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	//canal := make(chan string)

	// Thread 2
	//go func() {
	//	canal <- "Hello there!"
	//}()

	//msg := <-canal

	//fmt.Println(msg)

	data := make(chan int)

	qtdWorkers := 200000

	// Thread 2
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
