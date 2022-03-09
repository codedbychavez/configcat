package controllers

import (
	"fmt"

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

	// TODO: Do error checking on the width and height
	width := dimension.Width
	height := dimension.Height


	message := fmt.Sprintf("Your Width is: %.2f and the Height is: %.2f", width, height)

	return c.SendString(message)
}