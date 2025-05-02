package users

import (
	"context"

	usersmodelsv1 "users_service/internal/models/users/v1"
)

func (r *repository) AuthenticateUser(ctx context.Context, user usersmodelsv1.UserAuthenticateRequest) (*usersmodelsv1.UserAuthenticateResponse, error) {
	// var hashedPassword string
	// var userAuthenticateResponse models.UserAuthenticateResponse

	// row := r.db.DB().QueryRow(
	// 	ctx, `
	// 		SELECT id, password_hash, role
	// 		FROM users
	// 		WHERE email = $1;
	// 		`,
	// 	user.Login)

	// err := row.Scan(&userAuthenticateResponse.ID, &hashedPassword, &userAuthenticateResponse.Role)
	// if err != nil {
	// 	println(user.Login)
	// 	log.Printf("ошибка при аутентификации пользователя: %v", err)
	// 	if err == pgx.ErrNoRows {
	// 		return nil, errors.ErrUnauthorized
	// 	}
	// 	return nil, errors.ErrInternalError
	// }

	// if !hasher.CheckPasswordHash(user.Password, hashedPassword) {
	// 	log.Printf("неверный пароль для пользователя: %v", user.Login)
	// 	return nil, errors.ErrUnauthorized
	// }

	// return &userAuthenticateResponse, nil
	return nil, nil
}
