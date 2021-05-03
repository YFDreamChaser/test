package main

import (
	"fmt"
	"playermove"
)

func main() {
	p := playermove.NewPlayer(0.5)
	p.MoveTo(playermove.Vec2{X: 3, Y: 1})
	for !p.IsArrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
}
