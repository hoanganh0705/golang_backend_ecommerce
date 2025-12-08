package service

import "GolangBackendEcommerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service url
// controller -> service -> repo -> models -> database
func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUser()
}
