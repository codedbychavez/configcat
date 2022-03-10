package paint

type Colors struct {
	Color1 string
	Color2 string
}

func NewColors(color1 string, color2 string) Colors {
	return Colors{
		Color1: color1,
		Color2: color2,
	}
}                               

func Mix(color1 string, color2 string) string {
	mixResult := generateColor(color1, color2)
	return mixResult
}

func generateColor(color1 string, color2 string) string {
	var mixResult string
	colors := map[string]Colors{
		"orange": NewColors("red", "yellow"),
		"green" : NewColors("yellow", "blue"),
		"violet": NewColors("blue", "red"),
	}
	// Loop over colors map
	for result, colorsGroup := range colors {
		mixExpression1 := colorsGroup.Color1 == color1 && colorsGroup.Color2 == color2
		mixExpression2 := colorsGroup.Color1 == color2 && colorsGroup.Color2 == color1

		if mixExpression1 || mixExpression2 {
			mixResult = result
			break
		}
	}

	if mixResult == "" {
		mixResult = "Our library does not support mixing these two colors together."
	}

	return mixResult
}