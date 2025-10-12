package auth

import (
	"errors"
	"unicode"
)

func ValidatePassword(password string) error {
	var (
		hasMinLen = len(password) >= 8
		hasUpper  = false
		hasLower  = false
		hasDigit  = false
	)

	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		}
	}

	switch {
	case !hasMinLen:
		return errors.New("password must be at least 8 characters")
	case !hasUpper:
		return errors.New("password must contain an uppercase letter")
	case !hasLower:
		return errors.New("password must contain a lowercase letter")
	case !hasDigit:
		return errors.New("password must contain a digit")
	default:
		return nil
	}
}
