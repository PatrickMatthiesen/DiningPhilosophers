package main

import (
	"fmt"
	"time"
)

const rounds = 50 // number of times a philospher must eat
const eat_time = 100 // eating time

func philosopher(Left, Right chan string, id int) {

	counter := 0

	for i := 0; i < rounds; i++ {
		Left <- "acquire"
		<-Left
		Right <- "acquire"
		<-Right

		fmt.Println("Hello, I am philosopher", id, "and have already eaten", counter, "times.")
		time.Sleep(eat_time) // Eating...

		Left <- "release"
		Right <- "release"
		counter++
	}

	fmt.Println("philosopher", id, " terminated!")
}

func fork(Left, Right chan string) {
	for {
		select {
		case <-Left:
			Left <- "granted"
			<-Left
		case <-Right:
			Right <- "granted"
			<-Right
		}
	}
}

func main() {
	fmt.Println("Hi, this is the dining philosophers problem")

	// each fork has two channels - one for receiving requests and one for replying
	fork0L, fork0R := make(chan string), make(chan string)
	fork1L, fork1R := make(chan string), make(chan string)
	fork2L, fork2R := make(chan string), make(chan string)
	fork3L, fork3R := make(chan string), make(chan string)
	fork4L, fork4R := make(chan string), make(chan string)

	go fork(fork0L, fork0R)
	go fork(fork1L, fork1R)
	go fork(fork2L, fork2R)
	go fork(fork3L, fork3R)
	go fork(fork4L, fork4R)

	go philosopher(fork0R, fork1L, 0)
	go philosopher(fork1R, fork2L, 1)
	go philosopher(fork2R, fork3L, 2)
	go philosopher(fork3R, fork4L, 3)
	//	go philosopher(fork4OutR, fork4InR, fork0OutL, fork0InL, query4, 4) // with possible deadlock
	go philosopher(fork0L, fork4R, 4) // no deadlock

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
