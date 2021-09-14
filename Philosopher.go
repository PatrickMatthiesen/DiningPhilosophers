package main

import (
	"fmt"
	"math/rand"
	"time"
)

type philosopher struct {
	forkLeft, forkRight *fork
	id                  int
	timesEaten          int
}

func NewPhil(id int, left fork, right fork) {
	p := new(philosopher)
	p.id = id
	p.forkLeft = &left
	p.forkRight = &right

	go getForks(p)
}

func (p philosopher) eat() {
	p.timesEaten++
	fmt.Printf("Philosopher %d is eating for the %d. time\n", p.id, p.timesEaten)
}

func getForks(p *philosopher) {
	for {
		select {
		case <-p.forkLeft.free:
			{
				select {
				case <-p.forkRight.free:
					{
						p.eat()
						p.forkLeft.used <- true
						p.forkRight.used <- true
						p.forkRight.free <- true
						think()
					}
				default:
				}
				p.forkLeft.free <- true
			}
		default:
		}
	}

}

func think() {
	time.Sleep(time.Duration(rand.Int31n(2) * int32(time.Second)))
}

/*
the folowing methods might work after making the channels buffered -_-

func getForks(p *philosopher)  {
	fmt.Printf("phil %d tries to get forks\n", p.id)
	if getFork(p.forkLeft){
		fmt.Printf("Philosopher %d picked up fork %d\n", p.id, p.forkLeft.id)
		if getFork(p.forkRight) {
			fmt.Printf("Philosopher %d picked up fork %d\n", p.id, p.forkRight.id)
			p.eat()
			p.forkRight.fChIn <- "im done"
			fmt.Printf("Philosopher %d returns fork %d\n", p.id, p.forkRight.id)
		}
		fmt.Printf("Philosopher %d returns fork %d\n", p.id, p.forkLeft.id)
		p.forkLeft.fChIn <- "im done"
	}

	think(p)
}

func getFork(f *fork) bool {
	fmt.Printf("tries to get fork %d\n", f.id)
	f.fChIn <- "fork ready?"
	var answer string = <-f.fChOut
	fmt.Println(answer)
	return answer == "avaliable"
} */
