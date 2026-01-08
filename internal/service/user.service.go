package service

import (
	"GolangBackendEcommerce/internal/repo"
	"GolangBackendEcommerce/internal/utils/random"
	"GolangBackendEcommerce/pkg/response"
	"fmt"
)

type IUserService interface {
	Register(email string, purpose string) int
}

// ở đây không cần dùng pointer vì interface đã là 1 con trỏ ẩn rồi
type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email

	// 5. check OTP is available

	// 6. user spam...

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email, purpose) {
		return response.ErrCodeUserHasExist
	}

	// 2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("otp is :::%d\n", otp)

	// 3. Save OTP to redis with expiration time

	// 4. Send OTP to email

	// check email exists or not

	return response.ErrCodeSuccess
}
