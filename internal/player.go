package internal

import (
	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type player struct {
	pos          rl.Vector2
	vel          rl.Vector2
	standing     bool
	acceleration float32
	stopSpeed    float32
	jumpForce    float32
}

func newPlayer() *player {
	return &player{
		pos:          rl.NewVector2(screenWidth/2, screenHeight-playerHeight),
		vel:          rl.NewVector2(0, 0),
		standing:     true,
		acceleration: playerAcceleration,
		stopSpeed:    playerStopSpeed,
		jumpForce:    playerJumpForce,
	}
}

func (p *player) Draw() {
	rl.DrawRectangle(int32(p.pos.X), int32(p.pos.Y), playerWidth, playerHeight, rl.Green)
}

func (p *player) Update(state *gameState, dt float32) {
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

	if p.pos.Y >= screenHeight-playerHeight {
		p.pos.Y = screenHeight - playerHeight
		p.vel.Y = 0
		p.standing = true
	}

	if p.pos.X < float32(screenWidthMargin) {
		p.pos.X = float32(screenWidthMargin)
		p.vel.X *= -1.5
		p.vel.Y *= 1.2
	}

	if p.pos.X > screenWidth-playerWidth-screenWidthMargin {
		p.pos.X = screenWidth - playerWidth - screenWidthMargin
		p.vel.X *= -1.5
		p.vel.Y *= 1.2
	}

}
