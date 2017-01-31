package grifts

import (
	"github.com/bketelsen/sciencemingle/models"
	. "github.com/markbates/grift/grift"
)

var _ = Add("seed-users", func(c *Context) error {

	bob := models.User{}
	bob.FirstName = "Bob"
	bob.LastName = "Jones"
	bob.Location = "Tampa, FL"
	bob.Gender = models.Male
	bob.PictureURL = "http://images.sk-static.com/images/media/profile_images/artists/378080/huge_avatar"

	err := models.DB.Create(&bob)
	if err != nil {
		return err
	}

	susan := models.User{}
	susan.FirstName = "Susan"
	susan.LastName = "Smith"
	susan.Location = "Tampa, FL"
	susan.Gender = models.Female
	susan.PictureURL = "https://pbs.twimg.com/profile_images/561961243469025280/PXHEshSH.jpeg"

	err = models.DB.Create(&susan)

	return nil
})
