package routing

import (
	"configcat-homework/app/controllers"
	// "configcat-homework/internal/services"

	"github.com/gofiber/fiber/v2"
)


func CreatePublicRoutes(router fiber.Router, controllerManager controllers.ControllerManager) {
	router.Use(func(c *fiber.Ctx) error {
		c.Set("Allow", "GET, POST, PUT")
		return c.Next()
	})

	// A status route to check that the router is OK
	router.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	router.Post("/paint/liters", controllerManager.PaintController.ReturnLiters)
	router.Post("/paint/mix", controllerManager.PaintController.ReturnMix)
	router.Post("/paint/recommend/stores", controllerManager.PaintController.Companies)
	router.Post("/paint/recommend/color", controllerManager.PaintController.ReturnMix)





}