package main

type Position struct {
	X int
	Y int
}

type movingObject interface {
	updateVelocity()
	updatePosition()
	initialize()
	collide()
}
