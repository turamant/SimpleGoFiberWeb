package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const profileUnknown = "unknown"
var counter int64 = 0

func main()  {
	webApp := fiber.New()

	webApp.Get("/counter", func(c *fiber.Ctx) error {
		return c.SendString(strconv.FormatInt(counter, 10))
	})
	webApp.Post("/counter", func(c *fiber.Ctx) error {
		counter++
		return c.SendStatus(http.StatusOK)
	})

	webApp.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("The best school for Software Engineers")
	})

	webApp.Get("/courses", func(c *fiber.Ctx) error {
		return c.SendString("Java, Python, Go")
	})

	webApp.Get("/address" , func(c *fiber.Ctx) error {
		return c.SendString("Moscow, Red square 11a")
	})

    webApp.Get("/profiles", func(c *fiber.Ctx) error {
		profileID := c.Query("profile_id", profileUnknown)
		if profileID == "" {
			profileID = profileUnknown
		}
		if profileID == profileUnknown {
			return c.Status(http.StatusUnprocessableEntity).SendString("profile_id is required")
		}
		return c.SendString(fmt.Sprintf("user Profile ID: %s", profileID))
	})

	port := "4000"
	logrus.Fatal(webApp.Listen(":" + port))
	
	
}