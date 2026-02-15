package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	OBJECTMODE_DECORATION uint16 = iota
	OBJECTMODE_BLOCK
	OBJECTMODE_SPIKE
	OBJECTMODE_TRIGGER
	OBJECTMODE_PORTAL
	OBJECTMODE_PAD
	OBJECTMODE_ORB
)

var ObjectSprites []string = []string{"Resources/testBlock.png", "Resources/testSpike.png"}

type LevelObject struct {
	rectpro        RectPro
	colliderTop    RectPro
	colliderBottom RectPro
	id             uint
	mode           uint16

	sprite rl.Texture2D
	color  rl.Color
	depth  int8
}

func NewObject(rectPro RectPro, collider RectPro, id uint, mode uint16, depth int8) LevelObject {
	return LevelObject{
		rectpro:        rectPro,
		colliderTop:    collider,
		colliderBottom: collider,
		id:             id,
		mode:           mode,

		sprite: rl.LoadTexture(ObjectSprites[id-1]),
		color:  rl.White,
		depth:  depth,
	}
}

func NewBlock(rectpro RectPro, id uint, mode uint16, depth int8) LevelObject {
	return LevelObject{
		rectpro:        rectpro,
		colliderTop:    NewRectPro(rectpro.rect.X, rectpro.rect.Y-rectpro.origin.Y-1, rectpro.rect.Width, 2, 0),
		colliderBottom: NewRectPro(rectpro.rect.X, rectpro.rect.Y+rectpro.origin.Y+1, rectpro.rect.Width, 2, 0),
		id:             id,
		mode:           mode,

		sprite: rl.LoadTexture(ObjectSprites[id-1]),
		color:  rl.White,
		depth:  depth,
	}
}

func (object *LevelObject) Draw() {
	rl.DrawTextureEx(object.sprite, object.rectpro.GetPosition(), object.rectpro.rotation, 1, object.color)

	if *debug {
		rl.DrawRectangleLinesEx(object.rectpro.GetCollider(), 2, rl.Green)
		rl.DrawCircle(int32(object.rectpro.rect.X), int32(object.rectpro.rect.Y), 2, rl.Green)

		rl.DrawRectangleLinesEx(object.colliderTop.GetCollider(), 2, rl.Green)
		rl.DrawCircle(int32(object.colliderTop.rect.X), int32(object.colliderTop.rect.Y), 2, rl.Green)
	}
}
