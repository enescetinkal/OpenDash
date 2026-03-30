package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Editor struct {
	MousePosition rl.Vector2
}

func (s *Editor) Update(dt float32){
	s.MousePosition = rl.GetMousePosition()
	s.MousePosition = rl.GetWorldToScreen2D(s.MousePosition, mainCamera)

	if rl.IsKeyDown(rl.KeyD){
		mainCamera.Target.X++
	}
}

func InitalizeEditor() Editor {
	return Editor{
		MousePosition: rl.NewVector2(0, 0),
	}
}