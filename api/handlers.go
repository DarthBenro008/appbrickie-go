package api

import (
	"appbrickie/bot/handlers"
	"github.com/gofiber/fiber"
	"log"
	"strconv"
)

//Router
func HandlerRouter(app fiber.Router) {
	app.Get("/status", statusCheck)
	app.Get("/sendMessage/:id", sendID)
	app.Post("/sendPackage", sendPackage)
}

//HandlerFunctions
func statusCheck(c *fiber.Ctx) {
	_ = c.JSON(&fiber.Map{
		"msg":     handlers.HandlerBot.Self.UserName + " is up and running!",
		"success": true,
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

func sendPackage(c *fiber.Ctx) {
	file, err := c.FormFile("file")
	id := c.FormValue("id")
	msg := c.FormValue("msg")
	rid, _ := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handlers.SendErrorMessage(rid)
		log.Println("Invalid Multipart")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if file.Size > 50000000 {
		handlers.SendErrorMessage(rid)
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Exceeded Size Limit!",
		})
		return
	}
	filename := file.Filename
	err = c.SaveFile(file, "cache/"+file.Filename)
	if err != nil {
		handlers.SendErrorMessage(rid)
		log.Println("Error Saving File")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	response, err := handlers.SendPackage(rid, msg, "cache/"+file.Filename)
	if err != nil {
		handlers.SendErrorMessage(rid)
		_ = c.JSON(&fiber.Map{
			"success": false,
			"file":    filename,
			"error":   err.Error(),
		})
	}
	_ = c.JSON(&fiber.Map{
		"success": response,
		"file":    filename,
	})
}
