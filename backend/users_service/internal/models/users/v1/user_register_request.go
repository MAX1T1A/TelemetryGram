package users

import (
	"errors"
	"net/mail"
	"regexp"
	"slices"
)

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func IsValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

var roles = []string{"STUDENT", "TEACHER", "ADMIN"}

func IsValidRoleEnum(item string) bool {
	return slices.Contains(roles, item)
}

func (c *UserRegisterRequest) UserRegisterRequestValidate() error {
	if !IsValidUsername(c.Username) {
		return errors.New("неправильный формат ")
	}

	if !IsValidEmail(c.Email) {
		return errors.New("некорректный email")
	}

	if !IsValidRoleEnum(c.Role) {
		return errors.New("невереное значение role")
	}

	return nil
}
