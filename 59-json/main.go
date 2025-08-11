package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	cm1 := NewColourMaster(101, "red")

	bytes, err := json.Marshal(cm1)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(string(bytes))

	cm2 := &ColourMaster{}

	err = json.Unmarshal(bytes, cm2)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(cm2.Id, cm2.Code)
	}

	m1 := make(map[string]any)

	err = json.Unmarshal(bytes, &m1)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(m1)
	}

	var st1 struct {
		Id   uint   `json:"id"`
		Code string `json:"code"`
	}

	err = json.Unmarshal(bytes, &st1)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(st1.Code, st1.Id)
	}
}

func NewColourMaster(id uint, code string) *ColourMaster {
	return &ColourMaster{id, code}
}

type ColourMaster struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
}
