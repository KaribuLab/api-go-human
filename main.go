package main

import (
	"log"

	"github.com/Karibu/api-go-human/repository"
	"github.com/Karibu/api-go-human/route"
	"github.com/Karibu/api-go-human/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	userRepository, _ := repository.NewUserRepository()
	userService, _ := service.NewUserService(userRepository)
	app := fiber.New()
	route.NewUserRoute(app, userService)
	log.Fatal(app.Listen(":3000"))
}
