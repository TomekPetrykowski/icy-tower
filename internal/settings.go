package internal

import "github.com/chewxy/math32"

const playerWidth = 50
const playerHeight = 50
const screenWidth = 1280
const screenHeight = 720
const targetFps = 60
const playerAcceleration float32 = 5000
const levelFriction float32 = 4.5
const playerStopSpeed = playerAcceleration / 2
const playerJumpTime = 0.14 // lower - slower
const gravity float32 = 777 * screenHeight / 2 * (playerJumpTime * playerJumpTime)
const platformHeight float32 = 30
const screenWidthMargin = screenWidth / 7
const screenHeightMargin int32 = int32(platformHeight)
const platformScreenWidth int32 = screenWidth - 2*screenWidthMargin
const minPlatformWidth int32 = platformScreenWidth / 6
const maxPlatformWidth int32 = platformScreenWidth / 2

var playerJumpForce float32 = math32.Sqrt(2 * 777 * screenHeight * gravity)
