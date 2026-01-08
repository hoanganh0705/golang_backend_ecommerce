package service

import (
	"GolangBackendEcommerce/internal/repo"
	"GolangBackendEcommerce/internal/utils/crypto"
	"GolangBackendEcommerce/internal/utils/random"
	"GolangBackendEcommerce/internal/utils/sendto"
	"GolangBackendEcommerce/pkg/response"
	"fmt"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

// ở đây không cần dùng pointer vì interface đã là 1 con trỏ ẩn rồi
type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashed email is :::%s\n", hashEmail)

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
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. Send OTP to email
	// err = sendto.SendTextEmailOTP([]string{email}, "anh487303@gmail.com", strconv.Itoa(otp))
	err = sendto.SendTemplateEmailOTP([]string{email}, "anh487303@gmail.com", "otp-auth.html", map[string]interface{}{
		"Email": email,
		"OTP":   strconv.Itoa(otp),
	})
	if err != nil {
		return response.ErrSendEmailOTP
	}

	// check email exists or not

	return response.ErrCodeSuccess
}
