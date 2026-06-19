package internal

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameState struct {
	player    *player
	levels    []*level
	platforms []*platform

	gravity  float32
	friction float32
}

func newGameState() *gameState {
	level1 := newLevel(1, screenHeight, 3, 150, rl.Red)
	level2 := newLevel(5, level1.verticalEnd, 4, 100, rl.Green)

	// levels := []*level{level1}
	levels := []*level{level1, level2}
	platforms := make([]*platform, 0)

	for _, l := range levels {
		platforms = append(platforms, l.platforms...)
		for _, p := range platforms {
			print(fmt.Sprintf("%p ", p))
		}
		println()
	}

	return &gameState{
		player:    newPlayer(),
		levels:    levels,
		platforms: platforms,
		gravity:   gravity,
		friction:  levelFriction,
	}
}
