package main

type Ball struct {
	pos       Position
	Xvelocity int
	Yvelocity int
	width     int
	height    int
	art       byte
}

/*
updateVelocity()
updatePosition()
initialize()
collide()
*/

func (b *Ball) initialize() {
	b.Xvelocity = 1
	b.Yvelocity = 1
	b.pos.X = 30
	b.pos.Y = 10
	b.width = 3
	b.height = 3
}

func (b *Ball) collide() {

}

func (b *Ball) updateVelocity(x, y int) {
	b.Xvelocity = x
	b.Yvelocity = y
}

func (b *Ball) updatePosition() {
	b.pos.X += b.Xvelocity
	b.pos.Y += b.Yvelocity
}
