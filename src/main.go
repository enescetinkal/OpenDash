package main

import (
	"flag"
	"fmt"
	"log"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// cmd arguments
var debug *bool
var noSound *bool

var ObjectList []LevelObject
var ObjectSprites []rl.Texture2D = make([]rl.Texture2D, 2)

const (
	ScreenW, ScreenH  int32 = 800, 600
	LEVEL_OBJECTLIMIT int   = 32768
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

	ObjectList = []LevelObject{
		NewBlock(NewRectPro(0, 0, 64, 64, 0), 1, 0),
		NewSpike(NewRectPro(0, 0, 64, 64, 0), 2, 0),
	}
	
	ObjectSprites = []rl.Texture2D{rl.LoadTexture("Resources/testBlock.png"), rl.LoadTexture("Resources/testSpike.png")}

	var groundHeight float32 = float32(ScreenH) - 100
	groundRect := rl.NewRectangle(0, groundHeight, float32(ScreenW)*4, float32(ScreenH)*2)
	groundColor := rl.Blue

	player := NewPlayer(groundHeight)
	mainCamera := rl.NewCamera2D(rl.NewVector2(float32(ScreenH)-500, float32(ScreenW)/2), rl.NewVector2(player.rectpro.rect.X, 400), 0, 1)

	level, err := InitalizeLevel("idk.json")

	if err != nil {
		log.Fatal(err)
	}

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
			if *debug {
				print(player.onGround)
			}

			mainCamera.Target = rl.NewVector2(player.rectpro.rect.X, 400)

			groundRect.X = player.rectpro.rect.X - groundRect.Width/2

			if rl.IsKeyPressed(rl.KeyR) {
				player = NewPlayer(groundHeight)
			}

			if rl.IsKeyPressed(rl.KeyS) {
				err := SaveLevel("idk", level)
				if err != nil {
					log.Fatal(err)
				}
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
