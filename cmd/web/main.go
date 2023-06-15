package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const profileUnknown = "unknown"

func main()  {
	webApp := fiber.New()

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