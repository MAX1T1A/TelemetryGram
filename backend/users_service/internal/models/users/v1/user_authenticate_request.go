package users

import "errors"

type UserAuthenticateRequest struct {
	Login    string `json:"email"`
	Password string `json:"password"`
}

func (c *UserAuthenticateRequest) UserAuthenticateRequestValidate() error {
	if len(c.Password) == 0 || len(c.Password) < 8 {
		return errors.New("пароль не может быть пустой строкой и длинна должна быть больше/равна 8 символам")
	}

	return nil
}
