package main

import (
	"os"
	"encoding/json"
)

type PrefabObject struct{
	Id uint
	Width float32
	Height float32
	Type string
}

type PrefabList struct{
	Modes map[string]uint16
	Textures []string
	Objects []PrefabObject
}

func MakePrefabs() (PrefabList, error) {
	var tempPrefab PrefabList

	data, err := os.ReadFile("objectPreFabs.json")
	if err != nil {
		return PrefabList{}, err
	}
	err = json.Unmarshal(data, &tempPrefab)
	if err != nil {
		return PrefabList{}, err
	}

	return tempPrefab, nil
}