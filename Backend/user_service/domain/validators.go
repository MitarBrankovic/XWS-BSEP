package domain

import (
	"github.com/go-playground/validator"
	"regexp"
)

func NewUserValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidator)
	validate.RegisterValidation("name", nameValidator)
	return validate
}

func usernameValidator(fl validator.FieldLevel) bool {
	matchString, err := regexp.MatchString(`^[A-Za-z][A-Za-z0-9_]{4,19}$`, fl.Field().String())
	if err != nil {
		return false
	}
	return matchString
}

func nameValidator(fl validator.FieldLevel) bool {
	matchString, err := regexp.MatchString(`^[A-Z][a-z]+[\s]{1}[A-Z][a-z]+$`, fl.Field().String())
	if err != nil {
		return false
	}
	return matchString
}
