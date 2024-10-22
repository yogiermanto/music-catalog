package memberships

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"music-catalog/internal/models/memberships"
)

func (r *Repository) CreateUser(ctx context.Context, model *memberships.User) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *Repository) GetUser(ctx context.Context, email, username string, id int) (resp memberships.User, err error) {
	err = r.db.WithContext(ctx).
		Where("email = ?", email).
		Or("username = ?", username).
		Or("id = ?", id).
		First(&resp).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}
