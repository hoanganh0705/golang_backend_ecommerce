package repo

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/internal/database"
	"context"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user repo url
// func (ur *UserRepo) GetInfoUser() string {
// 	return "NguyenHoangAnh"
// }

// implicit implementation of interface
type IUserRepository interface {
	GetUserByEmail(email string, purpose string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string, purpose string) bool {
	ctx := context.Background()
	user, err := u.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
