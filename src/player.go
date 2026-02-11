package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//import "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	rectpro       RectPro
	blockCollider RectPro

	yVelocity float32
	jumpForce float32
	onGround  bool

	gravity float32
	speed   float32
	gameOver bool

	depth int8
}

func UpdatePlayer(self *Player, dt float32, groundHeight float32) {
	if self.gameOver {
		return
	}

	if self.rectpro.rect.Y+self.rectpro.rect.Height/2 >= groundHeight {
		self.rectpro.rect.Y = groundHeight - self.rectpro.rect.Height/2
		self.onGround = true
	}

	//apply gravity and jump force
	self.rectpro.rect.Y += self.yVelocity * dt
	if !self.onGround {
		self.yVelocity += self.gravity * dt
	} else {
		self.yVelocity = 0
	}

	self.rectpro.rect.X += self.speed * dt

	//jump
	if rl.IsKeyDown(rl.KeySpace) && self.onGround {
		self.yVelocity = self.jumpForce
	}

	self.blockCollider.rect.X = self.rectpro.rect.X
	self.blockCollider.rect.Y = self.rectpro.rect.Y
}

func UpdateCollisions(self *Player, object *LevelObject) {
	if(CheckCollisionRectPro(self.rectpro, object.colliderTop) && object.mode == OBJECTMODE_BLOCK){
		self.rectpro.rect.Y = object.rectpro.rect.Y - object.rectpro.origin.Y - self.rectpro.origin.Y
		self.blockCollider.rect.Y = object.rectpro.rect.Y - object.rectpro.origin.Y - self.rectpro.origin.Y
		self.onGround = true
	} else {
		self.onGround = false
	}

	if(CheckCollisionRectPro(self.blockCollider, object.rectpro) && object.mode == OBJECTMODE_BLOCK){
		self.gameOver = true
	}
}

func InitalizePlayer(groundHeight float32) Player {
	return Player{
		rectpro:       NewRectPro(0, float32(ScreenH)+groundHeight, 64, 64, 0),
		blockCollider: NewRectPro(0, float32(ScreenH)+groundHeight, 32, 32, 0),
		yVelocity:     0,
		jumpForce:     -500,
		gravity:       1000,
		speed:         350,
		onGround:      true,
		depth:         63,
	}
}
