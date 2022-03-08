package paint

import "fmt"

type Color struct {
	Color1 string
	Color2 string
}

func NewColor(color1 string, color2 string) Color {
	return Color{
		Color1: color1,
		Color2: color2,
	}
}                               

func Mix(color1 string, color2 string) string {
	generateColor("red", "blue")
	return "Mix Result"
}

func generateColor(color1 string, color2 string) {
	colors := map[string]Color{
		"orange": NewColor("red", "yellow"),
		"green" : NewColor("yellow", "blue"),
		"violet": NewColor("blue", "red"),
	}
	// Loop over colors map
	for key, value := range colors {
		fmt.Println(key, value)
	}
}