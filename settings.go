package main

import "github.com/chewxy/math32"

const PlayerWidth = 50
const PlayerHeight = 50
const ScreenWidth = 1280
const ScreenHeight = 720
const TargetFps = 60
const PlayerAcceleration float32 = 5000
const LevelFriction float32 = 4.5
const PlayerStopSpeed = PlayerAcceleration / 2
const PlayerJumpTime = 0.14 // lower - slower
const Gravity float32 = 777 * ScreenHeight / 2 * (PlayerJumpTime * PlayerJumpTime)

var PlayerJumpForce float32 = math32.Sqrt(2 * 777 * ScreenHeight * Gravity)
