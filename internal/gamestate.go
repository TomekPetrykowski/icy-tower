package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type gameState struct {
	player    *player
	platforms []*platform

	gravity  float32
	friction float32
}

func newGameState() *gameState {
	platforms := []*platform{
		{
			Pos:   rl.NewVector2(100, screenHeight-100),
			Size:  rl.NewVector2(300, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(700, screenHeight-200),
			Size:  rl.NewVector2(200, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(200, screenHeight-300),
			Size:  rl.NewVector2(250, 30),
			Color: rl.Red,
		},
	}

	return &gameState{
		player:    newPlayer(),
		platforms: platforms,
		gravity:   gravity,
		friction:  levelFriction,
	}
}
