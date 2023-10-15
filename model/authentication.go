package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Authentication struct {
	Email    string `bson:"email,omitempty" validate:"required,email"`
	Password string `bson:"password,omitempty" validate:"required,min=6,max=12"`
}

func (a *Authentication) Validate() []string {
	err := validator.New().Struct(a)
	if err != nil {
		var allErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			allErrors = append(allErrors, fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		}
		return allErrors
	}
	return nil
}
