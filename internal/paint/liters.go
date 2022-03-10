package paint

type Dimension struct {
	Width  float64 `json:"width" xml:"width" form:"width"`
	Height float64 `json:"height" xml:"height" form:"height"`
}

func Liters(width float64, height float64) float64 {
	area := width * height
	total := area / 10.0
	return total
}