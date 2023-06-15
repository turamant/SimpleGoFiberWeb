package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const profileUnknown = "unknown"
const requestParamkeyEvent = "event"

var counter int64 = 0
var counters = make(map[string]int64)

type (
	CreateLogEntryRequest struct{
		Message string 	`json: "message"`
		Level string   	`json: "level`
		Timestamp int64	`json: "timestamp`
	}

	CreateLogEntryResponse struct {
		ID string  	`json: "id"`
	}

	LogEntry struct {
		ID 		  string
		Message   string
		Level     string
		Timestamp int64
	}
)

var logs []LogEntry

func main()  {
	webApp := fiber.New()

	webApp.Post("/logs", func(c *fiber.Ctx) error {
		var request CreateLogEntryRequest
		if err := c.BodyParser(&request); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}
		logEntry := LogEntry{
			ID: uuid.New().String(),
			Message: request.Message,
			Level: request.Level,
			Timestamp: request.Timestamp,
		}
		logs = append(logs, logEntry)
		return c.JSON(CreateLogEntryResponse{
			ID: logEntry.ID,
		})
	})


	webApp.Get("/counter", func(c *fiber.Ctx) error {
		return c.SendString(strconv.FormatInt(counter, 10))
	})
	webApp.Post("/counter", func(c *fiber.Ctx) error {
		counter++
		return c.SendStatus(http.StatusOK)
	})

	webApp.Get("/counter/:event", func(c *fiber.Ctx) error {
		event := c.Params(requestParamkeyEvent, "")
		if event == "" {
			return c.SendStatus(http.StatusUnprocessableEntity)
		}
		eventCounter, ok := counters[event]
		if !ok {
			return c.SendStatus(http.StatusNotFound)
		}
		return c.SendString(strconv.FormatInt(eventCounter, 10))
	})

	webApp.Post("/counter/:event", func(c *fiber.Ctx) error {
		event := c.Params(requestParamkeyEvent, "")
		if event == "" {
			return c.SendStatus(http.StatusUnprocessableEntity)
		}
		counters[event] += 1
		
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