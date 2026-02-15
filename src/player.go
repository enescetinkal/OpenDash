package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	rectpro       RectPro
	blockCollider RectPro

	yVelocity float32
	jumpForce float32
	onGround  bool

	gravity float32
	speed   float32
	isDead  bool

	depth int8
}

func (p *Player) Update(dt float32, groundHeight float32) {
	if p.isDead {
		return
	}

	if rl.IsKeyDown(rl.KeySpace) && p.onGround {
		p.yVelocity = p.jumpForce
	}

	if p.rectpro.rect.Y+p.rectpro.origin.Y >= groundHeight {
		p.rectpro.rect.Y = groundHeight - p.rectpro.origin.Y
		p.onGround = true
	} else {
		p.onGround = false
	}

	p.rectpro.rect.Y += p.yVelocity * dt
	if !p.onGround {
		p.yVelocity += p.gravity * dt
	} else {
		p.yVelocity = 0
	}

	p.rectpro.rect.X += p.speed * dt

	p.blockCollider.rect.X = p.rectpro.rect.X
	p.blockCollider.rect.Y = p.rectpro.rect.Y

}

func (p *Player) UpdateCollisions(object *LevelObject) {
	if p.rectpro.CheckCollision(object.colliders[0]) && object.mode == OBJECTMODE_BLOCK {
		p.rectpro.rect.Y = object.rectpro.rect.Y - object.rectpro.origin.Y - p.rectpro.origin.Y
		p.blockCollider.rect.Y = object.rectpro.rect.Y - object.rectpro.origin.Y - p.rectpro.origin.Y
		p.yVelocity = 0
		p.onGround = true
	}

	if p.blockCollider.CheckCollision(object.rectpro) && object.mode == OBJECTMODE_BLOCK {
		p.isDead = true
	}
}

func NewPlayer(groundHeight float32) Player {
	return Player{
		rectpro:       NewRectPro(0, float32(ScreenH)+groundHeight, 64, 64, 0),
		blockCollider: NewRectPro(0, float32(ScreenH)+groundHeight, 32, 32, 0),
		yVelocity:     0,
		jumpForce:     -500,
		gravity:       1200,
		speed:         350,
		onGround:      true,
		depth:         63,
	}
}
