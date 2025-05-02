package users

import (
	"context"

	usersmodelsv1 "users_service/internal/models/users/v1"
)

func (r *repository) CreateUser(ctx context.Context, user usersmodelsv1.UserRegisterRequest) error {
	// hashedPassword, err := hasher.HashPassword(newPassword)
	// if err != nil {
	// 	log.Printf("ошибка хеширования пароля: %v", err)
	// 	return errors.ErrInternalError
	// }

	// row := r.db.DB().QueryRow(
	// 	ctx, `
	// 		INSERT INTO users (surname, name, patronymic, email, password_hash, role)
	// 		VALUES ($1, $2, $3, $4, $5, $6)
	// 		RETURNING id;
	// 		`,
	// 	user.Surname, user.Name, user.Patronymic, user.Email, hashedPassword, user.Role)

	// var userID string
	// err = row.Scan(&userID)
	// if err != nil {
	// 	if pgErr, ok := err.(*pgconn.PgError); ok {
	// 		if pgErr.Code == "23505" {
	// 			log.Printf("ошибка: дублирующийся ключ для email %v", user.Email)
	// 			return errors.DuplicateWithDetails(
	// 				"Пользователь с таким email уже существует",
	// 			)
	// 		}
	// 	}

	// 	log.Printf("ошибка при добавлении пользователя: %v", err)
	// 	return errors.ErrInternalError
	// }

	return nil
}
