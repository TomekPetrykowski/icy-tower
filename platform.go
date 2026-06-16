package main

import (
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

func (p *Platform) CollideWith(player *Player) bool {
	// On top
	if player.Pos.X <= p.Pos.X+p.Size.X && player.Pos.X+PlayerWidth >= p.Pos.X {
		if math32.Abs((player.Pos.Y+PlayerHeight)-p.Pos.Y) < 10 && player.Velocity.Y >= 0 {
			return true
		}
	}
	return false
}
