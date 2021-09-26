package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"notify-server/internal/biz"
	"notify-server/internal/data/ent"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *ent.Client
}

func NewUserRepo(data *ent.Client) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ur *userRepo) Create(ctx context.Context, user biz.User) (*biz.User, error) {
	uuid := uuid.New()
	u, err := ur.data.User.Create().SetUUID(uuid).SetEmail(user.Email).
		SetMobile(user.Mobile).SetName(user.Name).Save(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "创建用户失败")
	}
	result := new(biz.User)
	result.Id = u.ID
	result.Name = u.Name
	result.Email = u.Email
	result.Mobile = u.Mobile
	result.CreatedAt = u.CreatedAt
	result.UpdatedAt = u.UpdatedAt
	result.Status = int8(u.Status)
	return result, nil
}
