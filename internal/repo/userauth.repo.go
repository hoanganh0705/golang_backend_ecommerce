package repo

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type UserAuthRepository struct {
}

// AddOTP implements IUserAuthRepository.
func (u *UserAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	panic("unimplemented")
}

// constructor function
func NewUserAuthRepository() IUserAuthRepository {
	return &UserAuthRepository{}
}
