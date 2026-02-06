package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Mode Enum
const (
	OBJECTMODE_DECORATION uint16 = 0
	OBJECTMODE_BLOCK uint16 = 1
	OBJECTMODE_SPIKE uint16 = 2 // rectangular hitbox
)

type LevelObject struct {
	rectpro RectPro

	id uint
	mode uint16
	sprite rl.Texture2D
}