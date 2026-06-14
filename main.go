// This package contains game loop for icy-tower game
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1200, 600, "Hello World!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Hello World!", 10, 10, 25, rl.White)

		rl.EndDrawing()
	}
}
