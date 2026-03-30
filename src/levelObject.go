package main

import (
	"slices"
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

type LevelObject struct {
	rectpro   RectPro
	colliders []RectPro
	id        uint
	mode      uint16

	sprite rl.Texture2D
	color  rl.Color
	depth  int8
}

func NewObject(rectPro RectPro, id uint, mode uint16, depth int8) LevelObject {
	tempObject := LevelObject{
		rectpro:   rectPro,
		colliders: []RectPro{},
		id:        id,
		mode:      mode,

		sprite: ObjectSprites[id-1],
		color:  rl.White,
		depth:  depth,
	}

	tempObject.CreateColliders()

	return tempObject
}

// Makes a LevelObject from a CondensedObject.
func NewObjectFromReference(objList []LevelObject, condenced CondensedObject) LevelObject {
	// yes, i do know that "objList []LevelObject" is taxing on the memory
	tempObject := objList[condenced.Id-1]

	// thanks @NULLDEREF for this fix
	tempObject.colliders = slices.Clone(tempObject.colliders)

	tempObject.rectpro.rect.X = condenced.X
	tempObject.rectpro.rect.Y = condenced.Y
	tempObject.rectpro.Rotate(condenced.Rotation)

	tempObject.CreateColliders()

	tempObject.depth = condenced.Depth

	return tempObject
}

func (object *LevelObject) Condence() CondensedObject {
	return CondensedObject{
		X:        object.rectpro.rect.X,
		Y:        object.rectpro.rect.Y,
		Rotation: object.rectpro.rotation,
		Id:       object.id,
		Depth:    object.depth,
	}
}

func (object *LevelObject) CreateColliders() {
	switch(object.mode){
	case OBJECTMODE_BLOCK:
		object.colliders = []RectPro{
			NewRectPro(object.rectpro.rect.X, object.rectpro.rect.Y-object.rectpro.origin.Y-1, object.rectpro.rect.Width-16, 2, 0), 
			NewRectPro(object.rectpro.rect.X, object.rectpro.rect.Y+object.rectpro.origin.Y+1, object.rectpro.rect.Width-16, 2, 0),
		}
	case OBJECTMODE_SPIKE:
		object.colliders = []RectPro{NewRectPro(object.rectpro.rect.X, object.rectpro.rect.Y, object.rectpro.rect.Width/4, object.rectpro.rect.Height/2, object.rectpro.rotation)}
	}
}

func NewBlock(rectpro RectPro, id uint, depth int8) LevelObject {
	return NewObject(rectpro, id, OBJECTMODE_BLOCK, depth)
}

func NewSpike(rectpro RectPro, id uint, depth int8) LevelObject {
	return NewObject(rectpro, id, OBJECTMODE_SPIKE, depth)
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

func (object *LevelObject) IsValid() bool {
	return object.id > 0 && object.id <= uint(len(ObjectList))
}
