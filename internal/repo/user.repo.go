package repo

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/internal/model"
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

type UserRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *UserRepository) GetUserByEmail(email string, purpose string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
