package main

import (
	"fmt"

	"time"
)

func World(done chan int) {
	fmt.Println("World")
	done <- 1
}

func Ticker(seconds int, tick chan int) {
	for {
		time.Sleep(time.Duration(seconds) * time.Second)
		tick <- 1
	}
}

func main() {
	fmt.Println("Hello")
	tick := make(chan int)
	tock := make(chan int)

	go Ticker(2, tick)
	time.Sleep(time.Second)
	go Ticker(2, tock)

	for {
		select {
		case <-tick:
			fmt.Printf("tick...")
		case <-tock:
			fmt.Printf("TOCK...")
		}
	}
}
