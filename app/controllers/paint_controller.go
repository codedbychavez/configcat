package controllers

import (
	"configcat-homework/internal/paint"

	"github.com/gofiber/fiber/v2"
)

type PaintController struct {}

type PaintControllerManager interface {
	ReturnLiters(c *fiber.Ctx) error
}


type Dimension struct {
	Width float64 `json:"width" xml:"width" form:"width"`
	Height float64 `json:"height" xml:"height" form:"height"`
}


func (ctrl PaintController) ReturnLiters(c *fiber.Ctx) error {

	dimension := new(Dimension)

	if err := c.BodyParser(dimension); err != nil {
		return c.SendString("There was an error")
	}

	width := dimension.Width
	height := dimension.Height
	
	totalLiters := paint.Liters(width, height)

	return c.JSON(totalLiters)
}