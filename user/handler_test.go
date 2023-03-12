package user

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"user-service/db"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler_ReturnsStatus200(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)

	repo := NewRepo(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/api/users/:id", handler.Get)

	id, err := repo.Create(Model{Name: "test", Email: "test@gmail.com"})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/users/%d", id), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, 200, res.StatusCode)
}
