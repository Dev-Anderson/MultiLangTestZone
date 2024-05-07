package utils

import (
	"api-cookies/handlers/models"
	"errors"
	"regexp"
)

type ValidationResult struct {
	IsValid bool
	Error   error
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewValidationResult(isValid bool, err error) ValidationResult {
	return ValidationResult{
		IsValid: isValid,
		Error:   err,
	}
}

func ValidateEmail(email string) ValidationResult {
	if email == "" {
		return NewValidationResult(false, errors.New("empty email is not allowed"))
	}

	if !emailRegex.MatchString(email) {
		return NewValidationResult(false, errors.New("email is not valid"))
	}
	return NewValidationResult(true, nil)
}

func ValidatePassword(password string) ValidationResult {
	if len(password) < 6 {
		return NewValidationResult(false, errors.New("min length of password should be 6 characters"))
	}
	return NewValidationResult(true, nil)
}

func ValidateUserAuth(userToAuth models.UserToAuth) []string {
	var erros []string
	if v := ValidateEmail(userToAuth.Email); !v.IsValid {
		erros = append(erros, v.Error.Error())
	}

	if v := ValidatePassword(userToAuth.Password); !v.IsValid {
		erros = append(erros, v.Error.Error())
	}
	return erros
}
