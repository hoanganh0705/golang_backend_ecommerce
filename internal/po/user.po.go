package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:uuid;type:varchar(255);not null;primary_key;unique;index:idx_uuid"`
	UserName string    `gorm:"column:user_name"`
	IsActive bool      `gorm:"column:is_active; type:boolean; default:true"`
	Role     []Role    `gorm:"many2many:go_user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
