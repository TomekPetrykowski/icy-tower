// This package contains game loop for icy-tower game
package main

import (
	"fmt"

	"github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Settings
	rl.InitWindow(ScreenWidth, ScreenHeight, "Icy Tower Remake v0.0.1")
	defer rl.CloseWindow()
	rl.SetTargetFPS(TargetFps)

	state := NewGameState()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// Events ---------------------------------------------------------------- //
		if rl.IsKeyDown(rl.KeyD) {
			state.player.vel.X += state.player.acceleration * dt
		}
		if rl.IsKeyDown(rl.KeyA) {
			state.player.vel.X -= state.player.acceleration * dt
		}
		if rl.IsKeyPressed(rl.KeySpace) && state.player.standing {
			// TODO: Clean this up, this is initial jumping, it depends on horizontal velocity
			state.player.vel.Y = (-state.player.jumpForce - (math32.Abs(state.player.vel.X)/dt)/4) * dt
			state.player.standing = false
		}

		// Update ---------------------------------------------------------------- //
		for _, platform := range state.platforms {
			if platform.IsCollidingWith(state.player) {
				state.player.standing = true
				state.player.pos.Y = platform.Pos.Y - PlayerHeight
				break
			}
			state.player.standing = false
		}

		state.player.Update(state, dt)

		// Drawing --------------------------------------------------------------- //
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		state.player.Draw()
		for _, platform := range state.platforms {
			platform.Draw()
		}
		rl.DrawText(fmt.Sprintf("Vel: %.4f, %.4f", state.player.vel.X, state.player.vel.Y), 10, 10, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Pos: %.4f, %.4f", state.player.pos.X, state.player.pos.Y), 10, 40, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Standing: %v", state.player.standing), 10, 70, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Jump force: %v", state.player.jumpForce), 10, 100, 21, rl.White)
		rl.DrawText(fmt.Sprintf("Gravity: %v", state.gravity), 10, 130, 21, rl.White)

		rl.EndDrawing()
	}
}
