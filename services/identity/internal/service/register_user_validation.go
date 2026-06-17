package service

import (
	"net/mail"
	"strings"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/dto"
)

func normalizeRegisterUserInput(input dto.RegisterUserInput) dto.RegisterUserInput {
	return dto.RegisterUserInput{
		Username: strings.TrimSpace(input.Username),
		Email:    strings.ToLower(strings.TrimSpace(input.Email)),
		Password: input.Password,
	}
}

func validateRegisterUserInput(input dto.RegisterUserInput) error {
	if input.Username == "" {
		return ErrInvalidUsername
	}

	if !isValidEmailFormat(input.Email) {
		return ErrInvalidEmail
	}

	if input.Password == "" {
		return ErrInvalidPassword
	}

	return nil
}

func isValidEmailFormat(email string) bool {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return address.Address == email
}
