package main

// import "fmt"

type fork struct {
	id     int
	timesUsed int
	inUse  bool
	free  chan bool
	used chan bool
	chanOut chan bool
}

func NewFork(id int) *fork {
	f := new(fork)
	f.id = id
	f.free = make(chan bool, 1)
	f.used = make(chan bool, 1)
	f.chanOut = make(chan bool, 1)
	f.timesUsed = 0
	f.inUse = true
	go run(f)
	return f
}

func run(f *fork) {
	f.free <- true
	for {
		select {
			case <-f.used: {
				f.timesUsed++
				f.chanOut<- true
			}
			default:
		}
	}
}
