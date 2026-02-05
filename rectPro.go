package main

import rl "github.com/gen2brain/raylib-go/raylib"

type RectPro struct {
	rect     rl.Rectangle
	origin   rl.Vector2
	rotation float32
}

func NewRectPro(x float32, y float32, w float32, h float32, r float32) RectPro {
	return RectPro{
		rect:     rl.NewRectangle(x, y, w, h),
		origin:   rl.NewVector2(w/2, h/2),
		rotation: r,
	}
}

func DrawRectPro(rectpro *RectPro, color rl.Color) {
	rl.DrawRectanglePro(rectpro.rect, rectpro.origin, rectpro.rotation, color)

	if *debug {
		rl.DrawRectangleLinesEx(GetRectProCollider(*rectpro), 2, rl.Green)
		rl.DrawCircle(int32(rectpro.rect.X), int32(rectpro.rect.Y), 2, rl.Green)
	}
}

func GetRectProCollider(rectpro RectPro) rl.Rectangle {
	return rl.NewRectangle(rectpro.rect.X-rectpro.rect.Width/2, rectpro.rect.Y-rectpro.rect.Height/2, rectpro.rect.Width, rectpro.rect.Height)
}

func GetRectProPosition(rectpro RectPro) rl.Vector2 {
	return rl.NewVector2(rectpro.rect.X, rectpro.rect.Y)
}

func CheckCollisionRectPro(recta RectPro, rectb RectPro) bool {
	return rl.CheckCollisionRecs(GetRectProCollider(recta), GetRectProCollider(rectb))
}
