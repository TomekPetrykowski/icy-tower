package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type level struct {
	platforms     []*platform
	color         rl.Color
	startNum      uint8
	quantity      uint8
	verticalGap   int32
	verticalStart float32
	verticalEnd   float32
}

func newLevel(startNum uint8, verticalStart float32, quantity uint8, gap int32, color rl.Color) *level {
	level := &level{
		color:         color,
		startNum:      startNum,
		quantity:      quantity,
		verticalGap:   gap,
		verticalStart: verticalStart,
		verticalEnd:   verticalStart - float32(quantity)*float32(gap),
	}
	level.generate()

	return level
}

func (l *level) draw() {
	for _, p := range l.platforms {
		p.Draw()
	}
}

func (l *level) generate() {
	l.platforms = make([]*platform, 0)

	var i int32
	for i = 0; i < int32(l.quantity); i++ {
		platformWidth := rl.GetRandomValue(minPlatformWidth, maxPlatformWidth)
		platformPosX := rl.GetRandomValue(screenWidthMargin, platformScreenWidth-platformWidth)

		platform := platform{
			Pos:    rl.NewVector2(float32(platformPosX), l.verticalStart-float32(screenHeightMargin+i*l.verticalGap)),
			Size:   rl.NewVector2(float32(platformWidth), platformHeight),
			Color:  l.color,
			number: l.startNum + uint8(i),
		}

		l.platforms = append(l.platforms, &platform)
	}
}
