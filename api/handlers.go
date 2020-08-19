package api

import (
	"appbrickie/bot/handlers"
	"github.com/gofiber/fiber"
	"log"
	"strconv"
)

//Router
func HandlerRouter(app fiber.Router) {
	app.Get("/status", StatusCheck)
	app.Get("/sendMessage/:id", sendID)
	app.Post("/sendPackage", sendPackage)
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

func sendPackage(c *fiber.Ctx) {
	file, err := c.FormFile("file")
	if file.Size > 50000000 {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "Exceeded Size Limit!",
		})
		return
	}
	if err != nil {
		log.Println("Invalid Multipart")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	filename := file.Filename
	id := c.FormValue("id")
	msg := c.FormValue("msg")
	rid, _ := strconv.ParseInt(id, 10, 64)
	err = c.SaveFile(file, "cache/"+file.Filename)
	if err != nil {
		log.Println("Error Saving File")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	response, err := handlers.SendPackage(rid, msg, "cache/"+file.Filename)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"file":    filename,
			"error":   err.Error(),
		})
	}
	c.JSON(&fiber.Map{
		"success": response,
		"file":    filename,
	})
}
