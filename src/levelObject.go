package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Mode Enum
const (// default hitbox dimensions can change depending on the object's ID
	OBJECTMODE_DECORATION uint16 = iota		// has no physical hitbox
	OBJECTMODE_BLOCK
	OBJECTMODE_SPIKE
	OBJECTMODE_TRIGGER										// has no physical hitbox
	OBJECTMODE_PORTAL
	OBJECTMODE_PAD
	OBJECTMODE_ORB
)

var ObjectSprites []string = []string{"Resources/testBlock.png"}


type LevelObject struct {
	rectpro RectPro
	id uint
	mode uint16

	sprite rl.Texture2D
	color rl.Color
	// which order the object gets drawn at (player depth = 63)
	depth int8
}

func NewObject(rectPro RectPro, id uint, mode uint16, depth int8) LevelObject{
	//TODO: make object init
	return LevelObject{
		rectpro: rectPro,
		id: id,
		mode: mode,

		sprite: rl.LoadTexture(ObjectSprites[id - 1]),
		color: rl.White,
		depth: depth,
	}
}

func DrawLevelObject(object *LevelObject) {
	rl.DrawTextureEx(object.sprite, GetRectProPosition(object.rectpro), object.rectpro.rotation, 1, object.color)

	if *debug {
		rl.DrawRectangleLinesEx(GetRectProCollider(object.rectpro), 2, rl.Green)
		rl.DrawCircle(int32(object.rectpro.rect.X), int32(object.rectpro.rect.Y), 2, rl.Green)
	}
}
