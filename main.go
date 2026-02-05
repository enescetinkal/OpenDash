package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug *bool

const ScreenW, ScreenH int32 = 800, 600

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(ScreenW, ScreenH, "Open-Dash")

	rl.SetTargetFPS(60)

	var dt float32
	debug = flag.Bool("debug", false, "Debug Mode")

	flag.Parse()

	backgroundColor := rl.SkyBlue

	var groundHeight float32 = float32(ScreenH) - 100
	groundRect := rl.NewRectangle(0, groundHeight, float32(ScreenW)*4, float32(ScreenH)*2)
	groundColor := rl.Blue

	player := InitalizePlayer(groundHeight)
	mainCamera := rl.NewCamera2D(rl.NewVector2(float32(ScreenH)-500, float32(ScreenW)/2), rl.NewVector2(player.rectpro.rect.X, 400), 0, 1)

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()

		UpdatePlayer(&player, dt, groundHeight)
		mainCamera.Target = rl.NewVector2(player.rectpro.rect.X, 400)

		groundRect.X = player.rectpro.rect.X - groundRect.Width / 2

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		rl.BeginMode2D(mainCamera)
		rl.DrawRectangleRec(groundRect, groundColor)
		DrawRectPro(&player.rectpro, rl.Green)
		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
