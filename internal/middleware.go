package internal

import (
	"github.com/tanishashrivas/goApi/internal/models"
)

func IsEmpty(c *models.Course) bool {
	return c.CourseName == ""
}
