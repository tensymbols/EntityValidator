package Validator

import (
	"errors"
	"strings"
)

func AdValidate(title string, text string) error {
	if len(title) >= 100 || title == "" {
		return errors.New("invalid ad title")
	}
	if len(text) >= 500 || text == "" {
		return errors.New("invalid ad text")
	}
	return nil
}
func UserValidate(nickname string, email string) error {
	if len(nickname) >= 100 || nickname == "" {
		return errors.New("invalid user nickname")
	}
	s := strings.Split(email, "@")
	if len(s) < 2 || s[0] == "" || s[1] == "" {
		return errors.New("invalid user email")
	}
	return nil
}
