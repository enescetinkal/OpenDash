package main

import (
	"flag"
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var debug *bool
var noSound *bool

const ScreenW, ScreenH int32 = 800, 600 //TODO: Read InitWindow sizes from a config file

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(ScreenW, ScreenH, "Open-Dash")

	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	var dt float32
	var exitWindow bool
	var showMessageBox bool

	debug = flag.Bool("debug", false, "Debug Mode")
	noSound = flag.Bool("noSound", false, "Disable Sounds")
	flag.Parse()

	backgroundColor := rl.SkyBlue

	var groundHeight float32 = float32(ScreenH) - 100
	groundRect := rl.NewRectangle(0, groundHeight, float32(ScreenW)*4, float32(ScreenH)*2)
	groundColor := rl.Blue

	player := NewPlayer(groundHeight)
	mainCamera := rl.NewCamera2D(rl.NewVector2(float32(ScreenH)-500, float32(ScreenW)/2), rl.NewVector2(player.rectpro.rect.X, 400), 0, 1)

	objects := make([]LevelObject, 5) //TODO: Write object initializer class
	objects[0] = NewBlock(NewRectPro(float32(ScreenW), float32(ScreenH)-100-32, 64, 64, 0), 1, OBJECTMODE_BLOCK, 0)
	objects[1] = NewBlock(NewRectPro(float32(ScreenW)+64, float32(ScreenH)-100-32, 64, 64, 0), 1, OBJECTMODE_BLOCK, 100)
	objects[2] = NewBlock(NewRectPro(float32(ScreenW)+128, float32(ScreenH)-100-32, 64, 64, 0), 1, OBJECTMODE_BLOCK, 100)

	for !exitWindow {
		if !showMessageBox {
			dt = rl.GetFrameTime()
			exitWindow = rl.WindowShouldClose()

			if rl.IsKeyPressed(rl.KeyEscape) {
				showMessageBox = !showMessageBox
			}

			player.Update(dt, groundHeight)
			for i := 0; i < len(objects); i++ {
				player.UpdateCollisions(&objects[i])
			}

			mainCamera.Target = rl.NewVector2(player.rectpro.rect.X, 400)

			groundRect.X = player.rectpro.rect.X - groundRect.Width/2

			if rl.IsKeyPressed(rl.KeyR) {
				player = NewPlayer(groundHeight)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		rl.BeginMode2D(mainCamera)
		rl.DrawRectangleRec(groundRect, groundColor)

		//object and player draw
		for i := int8(-127); i < 127; i++ {
			if player.depth == i {
				player.rectpro.Draw(rl.Lime)
				player.blockCollider.Draw(rl.Blue)
			}

			for j := int(0); j < len(objects); j++ {
				if objects[j].depth == i {
					objects[j].Draw()
				}
			}
		}

		rl.EndMode2D()

		if *debug {
			rl.DrawText(fmt.Sprintf("Y Velocity = %.2f", player.yVelocity), 10, 10, 24, rl.RayWhite)
		}

		if showMessageBox {
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(rl.Black, 0.4))
			var result int32 = gui.MessageBox(rl.NewRectangle(float32(rl.GetScreenWidth())/2-125, float32(rl.GetScreenHeight())/2-50, 250, 100), gui.IconText(gui.ICON_EXIT, "Close Window"), "Do you really want to exit?", "Yes;No")

			switch result {
			case 0, 2:
				showMessageBox = false
			case 1:
				exitWindow = true
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
