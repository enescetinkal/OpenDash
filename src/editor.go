package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Editor struct {
	MousePosition rl.Vector2
}

func (s *Editor) Update(dt float32){
	s.MousePosition.X = (s.MousePosition.X / mainCamera.Zoom) + mainCamera.Target.X - (float32(ScreenW))
	s.MousePosition.Y = (s.MousePosition.Y / mainCamera.Zoom) + mainCamera.Target.Y - (float32(ScreenH))
}

func InitalizeEditor() Editor {
	return Editor{
		MousePosition: rl.NewVector2(0, 0),
	}
}