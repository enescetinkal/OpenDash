package main

import (
	"encoding/json"
	"os"
	"time"
)

type CondensedObject struct {
	x        float32
	y        float32
	rotation float32
	id       uint
	depth    int8
}

type SavedLevel struct {
	Name       string
	UpdateDate time.Time
	Objects    []CondensedObject
}

type Level struct {
	name       string
	updateDate time.Time
	objects    []LevelObject
}

func InitalizeLevel(filename string) (Level, error) {
	var savedLevel SavedLevel

	data, err := os.ReadFile(filename)
	if err != nil {
		return Level{}, err
	}
	err = json.Unmarshal(data, &savedLevel)
	if err != nil {
		return Level{}, err
	}

	lvl := Level{
		name:       savedLevel.Name,
		updateDate: savedLevel.UpdateDate,
		objects:    make([]LevelObject, 0, len(savedLevel.Objects)),
	}

	for _, cond := range savedLevel.Objects {
		obj := NewObjectFromReference(ObjectList, cond)
		lvl.objects = append(lvl.objects, obj)
	}

	return lvl, nil
}

func SaveLevel(filename string, lvl Level) error {
	// Update the timestamp right before serializing
	lvl.updateDate = time.Now()

	savedLevel := SavedLevel{
		Name:       lvl.name,
		UpdateDate: lvl.updateDate,
		Objects:    make([]CondensedObject, 0, len(lvl.objects)),
	}

	for _, obj := range lvl.objects {
		if obj.IsValid() { // Only save objects with valid ids
			savedLevel.Objects = append(savedLevel.Objects, obj.Condence())
		}
	}

	data, err := json.MarshalIndent(savedLevel, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename+".json", data, 0644)
}
