package controllers

import (
	"configcat-homework/internal/custom_configcat"
	"configcat-homework/internal/paint"
	"configcat-homework/internal/services"

	"github.com/gofiber/fiber/v2"
)

type PaintController struct{}

type PaintControllerManager interface {
	ReturnLiters(c *fiber.Ctx) error
	Mix(c *fiber.Ctx) error
	Companies(c *fiber.Ctx) error
	RandomColor(c *fiber.Ctx) error
}

// Return the total liters given two dimensions
func (ctrl PaintController) ReturnLiters(c *fiber.Ctx) error {
	dimension := new(paint.Dimension)
	if err := c.BodyParser(dimension); err != nil {
		return c.SendString("There was an error")
	}
	width := dimension.Width
	height := dimension.Height
	totalLiters := paint.Liters(width, height)
	return c.JSON(totalLiters)
}

// Return a result of mixing two colors together
func (ctrl PaintController) Mix(c *fiber.Ctx) error {
	colors := new(paint.Colors)
	if err := c.BodyParser(colors); err != nil {
		return c.SendString("There was an error")
	}
	color1 := colors.Color1
	color2 := colors.Color2
	mixResult := paint.Mix(color1, color2)
	return c.JSON(mixResult)
}

// Return a list of all companies
func (ctrl PaintController) Companies(c *fiber.Ctx) error {
	companyService := services.CompanyService{}
	appConfigcat := custom_configcat.Client{}.Connect().Connection

	isCompaniesEnabled := appConfigcat.GetValue("companies", false)

	if isCompaniesEnabled == true {
		companies := companyService.FindAll()
		return c.JSON(companies)
	}
	
	return c.SendString("This feature is not available.")
}

// Return a random color
func (ctrl PaintController) RandomColor(c *fiber.Ctx) error {
	paintService := services.PaintService{}
	appConfigcat := custom_configcat.Client{}.Connect().Connection

	isRandomColorEnabled := appConfigcat.GetValue("randomColor", false)

	if isRandomColorEnabled == true {
		randomColor := paintService.RandomColor()
		return c.JSON(randomColor)
	}

	return c.SendString("This feature is not available.")
}


