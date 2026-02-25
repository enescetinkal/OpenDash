package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Editor struct {
	MousePosition rl.Vector2
}

func InitalizeEditor() Editor {
	return Editor{
		MousePosition: rl.NewVector2(0, 0),
	}
}