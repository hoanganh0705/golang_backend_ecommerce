package repo

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

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string, purpose string) bool {
	panic("unimplemented")
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
