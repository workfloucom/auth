package orm

import (
	"context"

	"gorm.io/gorm"
	"workflou.com/auth/pkg/user"
)

type Users struct {
	DB *gorm.DB
}

func (r *Users) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	res := r.DB.First(&u, "email = ?", email)

	if res.Error != nil {
		return nil, user.ErrUserNotFound
	}

	return &u, nil
}
