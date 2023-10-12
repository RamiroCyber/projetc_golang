package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty" validate:"required,email"`
	FirstName string             `bson:"first_name,omitempty" validate:"required"`
	LastName  string             `bson:"last_name,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

func (u *User) Validate() []string {
	err := validator.New().Struct(u)
	if err != nil {
		var allErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			allErrors = append(allErrors, fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		}
		return allErrors
	}
	return nil
}
