package main

import (
	"flag"
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//cmd arguments
var debug *bool
var noSound *bool

const (
	ScreenW, ScreenH int32 = 800, 600
	LEVEL_OBJECTLIMIT int = 32768
) //TODO: Read InitWindow sizes from a config file

func main() {
	rl.InitWindow(ScreenW, ScreenH, "Open-Dash")
	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.SetTargetFPS(240)
	rl.SetExitKey(0)

	var dt float32
	var exitWindow bool
	var showMessageBox bool

	debug = flag.Bool("debug", false, "Debug Mode")
	noSound = flag.Bool("noSound", false, "Disable Sounds")
	flag.Parse()

	backgroundColor := rl.SkyBlue

	/*var ObjectList []LevelObject = []LevelObject{
		NewBlock(NewRectPro(0, 0, 64, 64, 0), 1, 0),
		NewSpike(NewRectPro(0, 0, 64, 64, 0), 2, 0),
	}*/

	var groundHeight float32 = float32(ScreenH) - 100
	groundRect := rl.NewRectangle(0, groundHeight, float32(ScreenW)*4, float32(ScreenH)*2)
	groundColor := rl.Blue

	player := NewPlayer(groundHeight)
	mainCamera := rl.NewCamera2D(rl.NewVector2(float32(ScreenH)-500, float32(ScreenW)/2), rl.NewVector2(player.rectpro.rect.X, 400), 0, 1)

	level := Level {
		name: "idk",
		objects: make([]LevelObject, 8, LEVEL_OBJECTLIMIT),
	} //TODO: Write object initializer class

	level.objects[0] = NewBlock(NewRectPro(float32(ScreenW), float32(ScreenH)-100-32, 64, 64, 0), 1, 0)
	level.objects[1] = NewBlock(NewRectPro(float32(ScreenW)+64, float32(ScreenH)-100-32, 64, 64, 90), 1, 100)
	level.objects[2] = NewBlock(NewRectPro(float32(ScreenW)+(64*2), float32(ScreenH)-100-32, 64, 64, 270), 1, 100)
	level.objects[3] = NewBlock(NewRectPro(float32(ScreenW)+(64*3), float32(ScreenH)-100-(32*5), 64, 64, 0), 1, 100)
	level.objects[4] = NewBlock(NewRectPro(float32(ScreenW)+(64*8), float32(ScreenH)-100-(32*7), 64, 64, 0), 1, 100)
	level.objects[5] = NewBlock(NewRectPro(float32(ScreenW)+(64*12), float32(ScreenH)-100-(32*9), 64, 64, 0), 1, 100)
	level.objects[6] = NewBlock(NewRectPro(float32(ScreenW)+(64*17), float32(ScreenH)-100-(32*11), 64, 64, 0), 1, 100)
	level.objects[7] = NewSpike(NewRectPro(float32(ScreenW)+(64*4), float32(ScreenH)-100-32, 64, 64, 0), 2, 20)

	for !exitWindow {
		if !showMessageBox {
			dt = rl.GetFrameTime()
			exitWindow = rl.WindowShouldClose()

			if rl.IsKeyPressed(rl.KeyEscape) {
				showMessageBox = !showMessageBox
			}

			player.Update(dt, groundHeight)
			for i := 0; i < len(level.objects); i++ {
				player.UpdateCollisions(&level.objects[i])
			}
			if *debug {print(player.onGround)}

			mainCamera.Target = rl.NewVector2(player.rectpro.rect.X, 400)

			groundRect.X = player.rectpro.rect.X - groundRect.Width/2

			if rl.IsKeyPressed(rl.KeyR) {
				player = NewPlayer(groundHeight)
			}

			if rl.IsKeyPressed(rl.KeyS) {
				SaveLevel("idk", level)
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

			for j := int(0); j < len(level.objects); j++ {
				if level.objects[j].depth == i {
					level.objects[j].Draw()
				}
			}
		}

		rl.EndMode2D()

		if *debug {
			rl.DrawText(fmt.Sprintf("Y Velocity = %.2f", player.yVelocity), 10, 10, 30, rl.RayWhite)
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
