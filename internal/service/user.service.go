package service

import "GolangBackendEcommerce/internal/repo"

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // user service url
// // controller -> service -> repo -> models -> database
// func (us *UserService) GetInfoUserService() string {
// 	return us.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

// ở đây không cần dùng pointer vì interface đã là 1 con trỏ ẩn rồi
type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (u *userService) Register(email string, purpose string) int {
	panic("unimplemented")
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
