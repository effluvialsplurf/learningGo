package main

type Brick struct {
	pos    Position
	width  int
	height int
	art    byte
}

// TODO: figure out how to dynamically 'place' the brick with regards to the render
func (b *Brick) initialize() {
	//b.pos.X =
	//b.pos.Y =
	b.width = 5
	b.height = 5
}

func (b *Brick) collide() {

}

func (b *Brick) kill() {
	// TODO: figure out how to kill the brick
}
