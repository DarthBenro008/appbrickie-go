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
	var filename string
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Invalid Multipart")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	file := form.File["file"][0]
	id := form.Value["id"][0]
	msg := form.Value["msg"][0]
	if file.Size >= 52428800 {
		log.Println("File Limit Exceeded")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   "File Size Limit Exceeded!",
		})
		return
	}
	filename = file.Filename
	err = c.SaveFile(file, "cache/"+file.Filename)
	if err != nil {
		log.Println("Error Saving File")
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("invalid conversion")
	}
	_, err = handlers.SendPackage(rid, msg, "cache/"+filename)
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"file":    filename,
			"error":   err.Error(),
		})
	}
	_ = c.JSON(&fiber.Map{
		"success": true,
		"file":    filename,
		"error":   "",
	})
}
