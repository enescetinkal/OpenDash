package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	rectpro   RectPro
	colliders []RectPro
	id        uint
	mode      uint16

	sprite rl.Texture2D
	color  rl.Color
	depth  int8
}

func NewObject(rectPro RectPro, collider RectPro, id uint, mode uint16, depth int8) LevelObject {
	return LevelObject{
		rectpro:   rectPro,
		colliders: []RectPro{collider},
		id:        id,
		mode:      mode,

		sprite: rl.LoadTexture(ObjectSprites[id-1]),
		color:  rl.White,
		depth:  depth,
	}
}


// Makes a LevelObject from a CondensedObject.
func NewObjectFromReference(objList []LevelObject, condenced CondensedObject) LevelObject {
	// yes, i do know that "objList []LevelObject" is taxing on the memory
	tempObject := objList[condenced.id - 1]

	tempObject.rectpro.rect.X = condenced.x
	tempObject.rectpro.rect.Y = condenced.x
	tempObject.rectpro.rotation = condenced.rotation
	tempObject.depth = condenced.depth

	return tempObject
}

func (object *LevelObject) Condence() CondensedObject {
	return CondensedObject {
		x: object.rectpro.rect.X,
		y: object.rectpro.rect.Y,
		rotation: object.rectpro.rotation,
		id: object.id,
		depth: object.depth,
	}
}

func NewBlock(rectpro RectPro, id uint, depth int8) LevelObject {
	return LevelObject{
		rectpro:   rectpro,
		colliders: []RectPro{NewRectPro(rectpro.rect.X, rectpro.rect.Y-rectpro.origin.Y-1, rectpro.rect.Width-16, 2, 0), NewRectPro(rectpro.rect.X, rectpro.rect.Y+rectpro.origin.Y+1, rectpro.rect.Width-16, 2, 0)},
		id:        id,
		mode:      OBJECTMODE_BLOCK,

		sprite: rl.LoadTexture(ObjectSprites[id-1]),
		color:  rl.White,
		depth:  depth,
	}
}

func NewSpike(rectpro RectPro, id uint, depth int8) LevelObject {
	return LevelObject{
		rectpro:   rectpro,
		colliders: []RectPro{NewRectPro(rectpro.rect.X, rectpro.rect.Y, rectpro.rect.Width/4, rectpro.rect.Height/2, rectpro.rotation)},
		id:        id,
		mode:      OBJECTMODE_SPIKE,

		sprite: rl.LoadTexture(ObjectSprites[id-1]),
		color:  rl.White,
		depth:  depth,
	}
}

func (object *LevelObject) Draw() {
	rl.DrawTexturePro(object.sprite, rl.NewRectangle(0, 0, object.rectpro.rect.Width, object.rectpro.rect.Height), object.rectpro.rect, object.rectpro.origin, object.rectpro.rotation, object.color)

	if *debug {
		rl.DrawRectangleLinesEx(object.rectpro.GetCollider(), 2, rl.Green)
		rl.DrawCircle(int32(object.rectpro.rect.X), int32(object.rectpro.rect.Y), 2, rl.Green)

		for i := 0; i < len(object.colliders); i++ {
			rl.DrawRectangleLinesEx(object.colliders[i].GetCollider(), 2, rl.Yellow)
			//rl.DrawCircle(int32(object.colliders[i].rect.X), int32(object.colliders[i].rect.Y), 2, rl.Green)
		}
	}
}
