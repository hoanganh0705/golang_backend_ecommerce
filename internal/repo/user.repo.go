package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// user repo url
func (ur *UserRepo) GetInfoUser() string {
	return "NguyenHoangAnh"
}
