package route

import (
	"github.com/Karibu/api-go-human/model"
	"github.com/Karibu/api-go-human/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserRoute(a *fiber.App, s service.IUserService) error {
	a.Get("/users/:id", getUserById(s))
	a.Get("/users", getAllUsers(s))
	a.Post("/users", create(s))
	return nil
}

func create(s service.IUserService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := new(model.User)
		err := c.BodyParser(user)
		if err != nil {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
			})
			return err
		}
		s.Save(*user)
		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"user":    user,
		})
		return nil
	}
}

func getUserById(s service.IUserService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, _ := s.GetById(id)
		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"user":    user,
		})
		return nil
	}
}

func getAllUsers(s service.IUserService) func(*fiber.Ctx) error {
	users, _ := s.GetAll()
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"users":   users,
		})
	}
}
