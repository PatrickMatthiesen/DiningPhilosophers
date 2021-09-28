package main

import (
	"math/rand"
	"time"
)

type philosopher struct {
	forkLeft, forkRight *fork
	id                  int
	timesEaten          int
	chanOut				chan int
	chanIn				chan bool
}

func NewPhil(id int, left fork, right fork) *philosopher {
	p := new(philosopher)
	p.id = id
	p.forkLeft = &left
	p.forkRight = &right
	p.chanOut = make(chan int, 2)
	p.chanIn = make(chan bool, 1)

	go getForks(p)
	return p
}

func (p *philosopher) eat() {
	p.timesEaten++
}

func getForks(p *philosopher) {
	for {
		select {
		case <-p.chanIn:
			{
				p.chanOut <- p.timesEaten
				p.chanOut <- p.id
		}
		default:
		}

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

