package models

import (
	"time"

	"encoding/json"

	"github.com/markbates/pop"
	"github.com/markbates/validate/validators"

	"github.com/markbates/validate"

	"github.com/satori/go.uuid"
)

const Male = "MALE"
const Female = "FEMALE"
const Other = "OTHER/Not Specified"

type User struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UserName   string    `json:"user_name" db:"user_name"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Gender     string    `json:"gender" db:"gender"`
	PictureURL string    `json:"picture_url" db:"picture_url"`
	Location   string    `json:"location" db:"location"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.Validate(

		&validators.StringIsPresent{Field: u.UserName, Name: "UserName"},

		&validators.StringIsPresent{Field: u.Password, Name: "Password"},

		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},

		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},

		&validators.StringIsPresent{Field: u.Location, Name: "Location"},
	), nil

}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (u *User) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
