package main

type philosopher struct{
	forkLeft, forkRight *fork
	id int
	eaten int
	pChIn chan int
	pChOut chan int
}

func NewPhil(id int, left fork, right fork)  {
	p := new(philosopher)
	p.id = id
	p.forkLeft = &left
	p.forkRight = &right

	eating(p)
}

func eating(p philosopher)  {
	can i do the fork
	:yes do the fork, (fork say noone can do the fork now)

	else think
	you can fork (input)
	pickup 2 forks
	lock 

	print 
	unlock 
	output noeating 
}
