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
	rl.SetTargetFPS(TargetFps)

	// Objects
	player := Player{
		Pos:      rl.NewVector2(ScreenWidth/2, ScreenHeight-PlayerHeight),
		Velocity: rl.NewVector2(0, 0),
		Standing: true,
	}

	platforms := []Platform{
		{
			Pos:   rl.NewVector2(100, ScreenHeight-150),
			Size:  rl.NewVector2(300, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(700, ScreenHeight-300),
			Size:  rl.NewVector2(200, 30),
			Color: rl.Red,
		},
		{
			Pos:   rl.NewVector2(300, ScreenHeight-450),
			Size:  rl.NewVector2(150, 30),
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
			if player.Velocity.Y >= 0 {
				player.Velocity.Y += 100 * gravity * dt
			}
		}
		if math32.Abs(player.Velocity.X) < stopSpeed*dt {
			player.Velocity.X = 0
		}

		for _, platform := range platforms {
			if platform.CollideWith(&player) {
				player.Pos.Y = platform.Pos.Y - PlayerHeight
				player.Standing = true
				break
			}
			player.Standing = false
		}
		player.Update(dt)

		// Drawing --------------------------------------------------------------- //
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		player.Draw()
		for _, platform := range platforms {
			platform.Draw()
		}
		rl.DrawText(fmt.Sprintf("Vel: %.4f, %.4f", player.Velocity.X, player.Velocity.Y), 10, 10, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Pos: %.4f, %.4f", player.Pos.X, player.Pos.Y), 10, 40, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Standing: %v", player.Standing), 10, 70, 21, rl.White)

		rl.EndDrawing()
	}
}
