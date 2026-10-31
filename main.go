package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from Fiber!"})
	})
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("ok\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop
		_ = app.Shutdown()
	}()

	// ":port" with tcp listens on every IPv4 and IPv6 address.
	if err := app.Listen(":"+port, fiber.ListenConfig{ListenerNetwork: fiber.NetworkTCP, DisableStartupMessage: true}); err != nil {
		log.Fatal(err)
	}
}
