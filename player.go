package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//import "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	rectpro RectPro

	yVelocity float32
	jumpForce float32
	onGround  bool

	gravity float32
	speed   float32
}

func UpdatePlayer(self *Player, dt float32, groundHeight float32) {
	if self.rectpro.rect.Y+self.rectpro.rect.Height/2 >= groundHeight {
		self.rectpro.rect.Y = groundHeight - self.rectpro.rect.Height/2
		self.onGround = true
	} else {
		self.onGround = false
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
		self.onGround = false
		self.yVelocity = self.jumpForce
	}
}

func InitalizePlayer(groundHeight float32) Player {
	return Player{
		rectpro:   NewRectPro(0, float32(ScreenH)+groundHeight+10, 64, 64, 0),
		yVelocity: 0,
		jumpForce: -500,
		gravity:   1000,
		speed:     350,
		onGround:  true,
	}
}
