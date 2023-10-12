package models

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty" validate:"required,email"`
	FirstName string             `bson:"first_name,omitempty" validate:"required"`
	LastName  string             `bson:"last_name,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required"`
}

func (u *User) Validate() error {
	err := validator.New().Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		}
	}
	return nil
}
