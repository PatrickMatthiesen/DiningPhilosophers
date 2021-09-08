package main

import()

func main() {
	p := new(Philosopher)

	p.n = 3
	go NewPhil()
	
}

func something(phil philosopher)  {
	
}