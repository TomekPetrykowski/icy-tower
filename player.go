package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Pos      rl.Vector2
	Velocity rl.Vector2
	Standing bool
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), PlayerWidth, PlayerHeight, rl.Green)
}

func (p *Player) Update(dt float32) {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Velocity, dt))

	if p.Pos.Y >= ScreenHeight-PlayerHeight {
		p.Pos.Y = ScreenHeight - PlayerHeight
		p.Standing = true
	}

	if p.Standing {
		p.Velocity.Y = 0
	}

	if p.Pos.X < 0 {
		p.Pos.X = 0
		p.Velocity.X *= -1.5
		p.Velocity.Y *= 1.2
	}

	if p.Pos.X > ScreenWidth-PlayerWidth {
		p.Pos.X = ScreenWidth - PlayerWidth
		p.Velocity.X *= -1.5
		p.Velocity.Y *= 1.2
	}

}
