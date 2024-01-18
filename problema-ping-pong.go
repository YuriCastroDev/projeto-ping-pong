package main

import (
	"fmt"
	"time"
)

func bater(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func rebater(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimr(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var c chan string = make(chan string)

	go bater(c)
	go imprimr(c)
	go rebater(c)

	var entrada string

	fmt.Scanln(&entrada)

}
