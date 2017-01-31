package actions

import (
	"errors"

	"github.com/bketelsen/sciencemingle/models"
	"github.com/gobuffalo/buffalo"
)

// UserShow default implementation.
func UserShow(c buffalo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.Error(400, errors.New("ID Required"))
	}
	var u models.User
	err := models.DB.Find(&u, id)
	if err != nil {
		return c.Error(500, err)

	}

	c.Set("u", u)

	return c.Render(200, r.HTML("user/show.html"))
}
