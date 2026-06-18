// Package main is an entrypoint to the game
package main

import (
	"fmt"

	"github.com/TomekPetrykowski/icy-tower/internal"
)

func main() {
	game := internal.NewGame()

	if err := game.Run(); err != nil {
		panic(fmt.Errorf("error: %w", err))
	}
}
