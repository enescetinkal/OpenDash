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
	color rl.Color
	// which order the object gets drawn at (player depth not determined yet)
	depth int8
}

func NewObject(){
	//TODO: make object init
}

func DrawLevelObject(object *LevelObject) {
	if *debug {
		rl.DrawRectangleLinesEx(GetRectProCollider(object.rectpro), 2, rl.Green)
		rl.DrawCircle(int32(object.rectpro.rect.X), int32(object.rectpro.rect.Y), 2, rl.Green)
	}

	rl.DrawTextureEx(object.sprite, GetRectProPosition(object.rectpro), object.rectpro.rotation, 1, object.color)
}