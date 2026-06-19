package internal

import (
	"fmt"

	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type platform struct {
	Pos    rl.Vector2
	Size   rl.Vector2
	Color  rl.Color
	number uint8
}

func (p *platform) Draw() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Size.X), int32(p.Size.Y), p.Color)
	text := fmt.Sprintf("%d", p.number)
	textWidth := rl.MeasureText(text, 20)
	rl.DrawText(
		text,
		int32(p.Pos.X)+(int32(p.Size.X)/2)-(textWidth/2),
		int32(p.Pos.Y)+(int32(p.Size.Y)/2)-10,
		20,
		rl.White,
	)
}

func (p *platform) IsCollidingWith(player *player) bool {
	// On top
	if player.pos.X <= p.Pos.X+p.Size.X && player.pos.X+playerWidth >= p.Pos.X {
		if math32.Abs((player.pos.Y+playerHeight)-p.Pos.Y) < 20 && player.vel.Y >= 0 {
			return true
		}
	}
	return false
}
