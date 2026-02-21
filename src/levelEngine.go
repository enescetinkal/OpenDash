package main

import (
	"os"
	"time"
	"encoding/json"
)

type CondensedObject struct{
	x float32
	y float32
	rotation float32
	id uint
	depth int8
}

type Level struct{
	name string
	updateDate time.Time
	objects []LevelObject
}

func InitalizeLevel(filename string, objList *[]LevelObject) (Level, error) {
    var lvl Level

    data, err := os.ReadFile(filename)
    if err != nil {
        return lvl, err
    }
    err = json.Unmarshal(data, &lvl)
    return lvl, err
}

func SaveLevel(filename string, lvl Level) error {
    // Update the timestamp right before serializing
    lvl.updateDate = time.Now()

    data, err := json.MarshalIndent(lvl, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}