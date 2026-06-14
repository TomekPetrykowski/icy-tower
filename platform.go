package main

import (
	"fmt"

	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Platform struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Color rl.Color
}

func (p *Platform) Draw() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Size.X), int32(p.Size.Y), p.Color)
}

func (p *Platform) CollideWith(player *Player) {
	// On top
	if player.Pos.X <= p.Pos.X+p.Size.X && player.Pos.X+PlayerWidth >= p.Pos.X {
		fmt.Printf("X good, Y difference %f \n", math32.Abs((player.Pos.Y+PlayerHeight)-p.Pos.Y))
		if math32.Abs((player.Pos.Y+PlayerHeight)-p.Pos.Y) < 10 && player.Velocity.Y > 0 {
			fmt.Printf("Y good\n")
			player.Pos.Y = p.Pos.Y - PlayerHeight
			player.Standing = true
		} else {
			player.Standing = false
		}
	} else {
		player.Standing = false
	}
}
