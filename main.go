// This package contains game loop for icy-tower game
package main

import (
	"fmt"

	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Settings
	// rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(ScreenWidth, ScreenHeight, "Hello World!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Objects
	player := Player{
		Pos:      rl.NewVector2(ScreenWidth/2, ScreenHeight-PlayerHeight),
		Velocity: rl.NewVector2(0, 0),
		Standing: true,
	}

	platforms := []Platform{
		{
			Pos:   rl.NewVector2(100, ScreenHeight-200),
			Size:  rl.NewVector2(300, 30),
			Color: rl.Red,
		},
	}

	var accel float32 = 5000
	var friction float32 = 4.5
	var gravity float32 = 9.7
	var stopSpeed = accel / 2

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// Events ---------------------------------------------------------------- //
		if rl.IsKeyDown(rl.KeyD) {
			player.Velocity.X += accel * dt
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.Velocity.X -= accel * dt
		}
		if rl.IsKeyPressed(rl.KeySpace) && player.Standing {
			player.Velocity.Y = accel * -10 * dt
			player.Standing = false
		}

		// Update ---------------------------------------------------------------- //

		player.Velocity.X -= player.Velocity.X * friction * dt
		if !player.Standing {
			player.Velocity.Y += 200 * gravity * dt
		}
		if math32.Abs(player.Velocity.X) < stopSpeed*dt {
			player.Velocity.X = 0
		}

		player.Pos = rl.Vector2Add(player.Pos, rl.Vector2Scale(player.Velocity, dt))

		for _, platform := range platforms {
			platform.CollideWith(&player)
		}
		player.Update()

		// Drawing --------------------------------------------------------------- //
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		player.Draw()
		for _, platform := range platforms {
			platform.Draw()
		}
		rl.DrawText(fmt.Sprintf("Vel: %.2f, %.2f", player.Velocity.X, player.Velocity.Y), 10, 10, 21, rl.White)

		rl.EndDrawing()
	}
}
