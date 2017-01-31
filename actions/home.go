package actions

import (
	"github.com/gobuffalo/buffalo"

	"github.com/bketelsen/sciencemingle/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	var users []*models.User
	err := models.DB.All(&users)
	if err != nil {
		c.Error(500, err)
	}

	c.Set("users", users)
	return c.Render(200, r.HTML("index.html"))
}
