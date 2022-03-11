package services

import (
	"math/rand"
	"time"
)

type PaintService struct {
}

type Color struct {
	Name string
}

type PaintMethods interface {
	RandomColor() string
}




func (paintService PaintService) RandomColor() string {
	colors := map[int]string{
		0: "red",
		1: "yellow",
		2: "green",
		3: "blue",
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	randomNumber := random.Intn(3)

	randomColor := colors[randomNumber]
	
	return randomColor
}