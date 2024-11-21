package main

type Paddle struct {
	pos       Position
	width     int
	Xvelocity int
	art       byte
}

/*
updateVelocity()
updatePosition()
initialize()
collide()
*/

func (p *Paddle) initialize() {
	p.Xvelocity = 0
	p.pos.X = 60
	p.pos.Y = 20
	p.width = 6
}

func (p *Paddle) collide() {

}

func (p *Paddle) updateVelocity(x int) {
	p.Xvelocity = 0
}

func (p *Paddle) updatePosition() {
	p.pos.X += p.Xvelocity
}
