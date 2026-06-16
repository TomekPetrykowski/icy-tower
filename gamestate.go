package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameState struct {
	player    *Player
	platforms []*Platform

	gravity  float32
	friction float32
}

func NewGameState() *GameState {
	platforms := []*Platform{
		{
			Pos:   rl.NewVector2(100, ScreenHeight-100),
			Size:  rl.NewVector2(300, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(700, ScreenHeight-200),
			Size:  rl.NewVector2(200, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(200, ScreenHeight-300),
			Size:  rl.NewVector2(250, 30),
			Color: rl.Red,
		},
	}

	return &GameState{
		player:    NewPlayer(),
		platforms: platforms,
		gravity:   Gravity,
		friction:  LevelFriction,
	}
}
