package repo

import (
	"GolangBackendEcommerce/global"
	"fmt"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type UserAuthRepository struct {
}

// AddOTP implements IUserAuthRepository.
func (u *UserAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // usr:<email>:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

// constructor function
func NewUserAuthRepository() IUserAuthRepository {
	return &UserAuthRepository{}
}
