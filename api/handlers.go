package api

import (
	"appbrickie/bot/handlers"
	"github.com/gofiber/fiber"
	"strconv"
)

//Router
func HandlerRouter(app fiber.Router) {
	app.Get("/status", StatusCheck)
	app.Get("/sendMessage/:id", sendID)
}

//HandlerFunctions
func StatusCheck(c *fiber.Ctx) {
	_ = c.JSON(&fiber.Map{
		"msg": "HelloWorld",
	})
}

func sendID(c *fiber.Ctx) {
	id := c.Params("id")
	realId, _ := strconv.ParseInt(id, 10, 64)
	res, err := handlers.SendMessage(realId, "This is an API automated Message!")
	_ = c.JSON(&fiber.Map{
		"success": res,
		"id":      id,
		"err":     err,
	})
}
