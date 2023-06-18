package main

import (
	"dockerGo/handler"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"log"
	"time"
)

func main() {

	//port := os.Getenv("PORT")
	//pport := ":" + port
	engine := html.New("./views", ".html")
	port := ":8080"
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:               1000,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString(" たからもの")
	//})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index - start with views directory
		return c.Render("inputtoken", fiber.Map{
			"Title":  "Update Bulk Token",
			"Status": true,
		})
	})

	app.Get("/ntpn/:ntpn", handler.GetNTPN)
	app.Post("/inserttoken", handler.InsertNTPN)
	app.Post("/ntpn", handler.BulkNTPN)
	app.Post("/insert", handler.InsertToken)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	s := gocron.NewScheduler(time.UTC)

	//cron job refresh token
	_, err := s.Every(5).Minutes().Do(handler.RefreshTokenUsingGetRequest)
	if err != nil {
		log.Println("Error refresh token every 5 minutes")
	}
	s.StartAsync()

	log.Fatal(app.Listen(port))

}
