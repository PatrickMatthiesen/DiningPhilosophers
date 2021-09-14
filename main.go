package main

import "fmt"

func main() {
	var forks [5]*fork
	fmt.Println("makes forks")
	for i := 0; i < 5; i++ {
		forks[i] = NewFork(i)
	}
	fmt.Println("makes phils")
	NewPhil(0, *forks[0], *forks[1])
	NewPhil(1, *forks[1], *forks[2])
	NewPhil(2, *forks[2], *forks[3])
	NewPhil(3, *forks[3], *forks[4])
	NewPhil(4, *forks[0], *forks[4])
/* 	for i := 0; i < 5; i++ {
		if i==4 {
			NewPhil(i,*forks[0], *forks[i])
			continue
		}
		NewPhil(i,*forks[i], *forks[i+1])
	}*/

	for {
		
	}
}
