package main

import "fmt"

type fork struct {
	id     int
	timesUsed int
	inUse  bool
	free  chan bool
	used chan bool
}

func NewFork(id int) *fork {
	f := new(fork)
	f.id = id
	f.free = make(chan bool, 1)
	f.used = make(chan bool, 1)
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
				fmt.Printf("fork %d has been used %d times\n",f.id, f.timesUsed)
			}
			default:
	
		}
	}
}

/* func run(f *fork) {
	//fmt.Printf("fork %d is waiting to recieve a request\n", f.id)
	input := <-f.fChIn
	//fmt.Printf("fork %d recieved %s\n", f.id, input)
	switch input {
	case "fork ready?":
		{
			if f.inUse {
				//fmt.Printf("fork %d is waiting for avaliable to be recieved\n", f.id)
				f.fChOut <- "avaliable"
				//fmt.Printf("fork %d avaliable is recieved\n", f.id)
				//time.Sleep(time.Second)
				f.inUse = false
				break
			}
			//fmt.Printf("fork %d is waiting for unavaliable to be recieved\n", f.id)
			f.fChOut <- "unavaliable"
			//fmt.Printf("fork %d unavaliable is recieved\n", f.id)
			//time.Sleep(time.Second)
			break
		}
	case "im done":
		{
			f.inUse = true
			break
		}
	}
} */
