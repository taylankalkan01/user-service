package main

import (
	"fmt"
	"log"
	"user-service/db"
	"user-service/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	fmt.Println(database.Config)

	repo := user.NewRepo(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal((err))
	}

	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/api/users/:id", handler.Get)
	app.Post("/api/users", handler.Create)
	app.Listen(":8080")
}
