package main

import (
	"fmt"
	"time"
)

func main() {
	var forks [5]*fork

	fmt.Println("makes forks")
	for i := 0; i < 5; i++ {
		forks[i] = NewFork(i)
	}

	fmt.Println("makes phils")

	var phils [5]*philosopher
	for i := 0; i < 5; i++ {
		if i==4 {
			phils[i] = NewPhil(i,*forks[0], *forks[i])
			continue
		}
		phils[i] = NewPhil(i,*forks[i], *forks[i+1])
	}

	for {
		for _, fork := range forks{
			fork.chanIn <- true
			select {
				case timesUsed := <- fork.chanOut: {
					fmt.Printf("fork %d has been used %d times\n", <-fork.chanOut, timesUsed)
				}
				default:
			}
		}
		for _, phil := range phils{
			phil.chanIn <- true
			select {
				case timesEaten := <- phil.chanOut: {
					fmt.Printf("Philosopher %d has eaten for the %d. time\n", <- phil.chanOut, timesEaten)
				}
				default:
			}
		}
		time.Sleep(time.Duration(2* int32(time.Second)))
	}
}

