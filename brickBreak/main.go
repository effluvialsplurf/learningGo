package main

import (
	"bytes"
	"fmt"
)

func main() {
	width := 60
	height := 20
	screen := make([][]byte, width)

	for h := 0; h < height; h++ {
		screen[h] = make([]byte, width)
	}

	for h := 0; h < height/4; h++ {
		for w := 0; w < width; w++ {
			screen[h][w] = 1
		}
	}

	buf := new(bytes.Buffer)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if screen[h][w] == 1 {
				buf.WriteByte('#')
			}
		}
		buf.WriteString("\n")
	}

	fmt.Println(buf.String())
}
