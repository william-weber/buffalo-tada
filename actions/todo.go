package actions

import (
	"net/http"
	"tada/models"

	"github.com/gobuffalo/buffalo"
)

// Lists all Todo items in order of creation
func TodoIndex(c buffalo.Context) error {
	todos := []models.Todo{}
	err := models.DB.Order("created_at desc").All(&todos)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(todos))
}

// Create a new Todo item with the specified text
func TodoAdd(c buffalo.Context) error {
	// Get text from request parameters
	text := c.Param("text")

	todo := models.Todo{Text: text}

	err := models.DB.Create(&todo)

	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(todo))
	}
	return c.Render(http.StatusOK, r.JSON(todo))
}
