package main

import (
	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos          rl.Vector2
	vel          rl.Vector2
	standing     bool
	acceleration float32
	stopSpeed    float32
	jumpForce    float32
}

func NewPlayer() *Player {
	return &Player{
		pos:          rl.NewVector2(ScreenWidth/2, ScreenHeight-PlayerHeight),
		vel:          rl.NewVector2(0, 0),
		standing:     true,
		acceleration: PlayerAcceleration,
		stopSpeed:    PlayerStopSpeed,
		jumpForce:    PlayerJumpForce,
	}
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.pos.X), int32(p.pos.Y), PlayerWidth, PlayerHeight, rl.Green)
}

func (p *Player) Update(state *GameState, dt float32) {
	p.vel.X -= p.vel.X * state.friction * dt

	if !p.standing {
		p.vel.Y += state.gravity * dt
		// if p.vel.Y > 0 {
		// 	p.vel.Y += state.gravity * dt
		// }
	}
	if math32.Abs(p.vel.X) < p.stopSpeed*dt {
		p.vel.X = 0
	}

	if p.standing {
		p.vel.Y = 0
	}

	p.pos = rl.Vector2Add(p.pos, rl.Vector2Scale(p.vel, dt))

	if p.pos.Y >= ScreenHeight-PlayerHeight {
		p.pos.Y = ScreenHeight - PlayerHeight
		p.standing = true
	}

	if p.pos.X < 0 {
		p.pos.X = 0
		p.vel.X *= -1.5
		p.vel.Y *= 1.2
	}

	if p.pos.X > ScreenWidth-PlayerWidth {
		p.pos.X = ScreenWidth - PlayerWidth
		p.vel.X *= -1.5
		p.vel.Y *= 1.2
	}

}
