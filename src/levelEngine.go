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

func InitalizeLevel(filename string) (Level, error) {
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
    jsonLevel := make([]CondensedObject, 8, LEVEL_OBJECTLIMIT)

    for i := range lvl.objects {
        cond := lvl.objects[i].Condence()
        jsonLevel = append(jsonLevel, cond)
    }


    data, err := json.MarshalIndent(jsonLevel, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename + ".json", data, 0644)
}