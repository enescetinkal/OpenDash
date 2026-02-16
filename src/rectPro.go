package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "math"

type RectPro struct {
	rect     rl.Rectangle
	origin   rl.Vector2
	rotation float32
}

func NewRectPro(x float32, y float32, w float32, h float32, r float32) RectPro {
	tempRectPro := RectPro{
		rect:     rl.NewRectangle(x, y, w, h),
		origin:   rl.NewVector2(w/2, h/2),
		rotation: 0,
	}
	tempRectPro.Rotate(r)
	return tempRectPro
}

func (rp *RectPro) Draw(color rl.Color) {
	rl.DrawRectanglePro(rp.rect, rp.origin, rp.rotation, color)

	if *debug {
		rl.DrawRectangleLinesEx(rp.GetCollider(), 2, rl.Green)
		rl.DrawCircle(int32(rp.rect.X), int32(rp.rect.Y), 2, rl.Green)
	}
}

func (rp RectPro) GetCollider() rl.Rectangle {
	return rl.NewRectangle(rp.rect.X-rp.rect.Width/2, rp.rect.Y-rp.rect.Height/2, rp.rect.Width, rp.rect.Height)
}

func (rp RectPro) GetPosition() rl.Vector2 {
	return rl.NewVector2(rp.rect.X-rp.rect.Width/2, rp.rect.Y-rp.rect.Height/2)
}

func (rp RectPro) CheckCollision(other RectPro) bool {
	return rl.CheckCollisionRecs(rp.GetCollider(), other.GetCollider())
}

func (rp RectPro) Rotate(r float32) {
	if int(math.Floor(float64(r) / 90)) % 2 == 0{
		tempWitdh := rp.rect.Width
		rp.rect.Width = rp.rect.Height
		rp.rect.Height = tempWitdh
	}
	
	rp.rotation = r
}
