package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty" validate:"required,email"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name"bson:"last_name,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required,min=6,max=12"`
	Phone     string             `bson:"phone,omitempty" validate:"required"`
	Role      string             `bson:"role,omitempty" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
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
