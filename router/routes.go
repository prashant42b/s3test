package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	s3 := app.Group("/s3")

	var router fiber.Router

	// query := router.Group("/lsBucketData")
	// query.Get("/:serialNumber", HandleSerialNumberQuerySearch)

	// search := router.Group("/search")
	// search.Post("/", routeHandler.HandleESSearch)

	//Setup node routers
	// trademarkRoutes.SetupTrademarkRoutes(api)

}
