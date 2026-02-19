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

func LoadLevel(filename string) (Level, error) {
    var lvl Level

    data, err := os.ReadFile(filename)
    if err != nil {
        return lvl, err
    }
    err = json.Unmarshal(data, &lvl)
    return lvl, err
}